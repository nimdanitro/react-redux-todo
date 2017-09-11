package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

// ProjectRepo contains projects
type ProjectRepo struct {
	db        *bolt.DB
	currentID uint64
}

// BucketName is the BoltDB bucket name for the projects
var BucketName = []byte("projects")

// NewRepo creates a new repository given the current path
func NewRepo(dbpath string) (*ProjectRepo, error) {

	db, err := bolt.Open(dbpath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(BucketName)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &ProjectRepo{db: db, currentID: 0}, nil
}

// Close the Repository
func (r *ProjectRepo) Close() {
	r.db.Close()
}

// Get a Project from the Repository by its ID
func (r *ProjectRepo) Get(id uint64) Project {
	var project Project
	err := r.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BucketName)
		if bucket == nil {
			return fmt.Errorf("Projects Bucket %q not found!", BucketName)
		}

		// make a byte key from the ID
		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, id)

		// get the Project by its ID key from the DB
		p := bucket.Get(key)
		err := json.Unmarshal(p, &project)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return Project{}
	}
	// return empty Projects list if not found
	return project
}

//GetAll returns a slice of project containign all projects
func (r *ProjectRepo) GetAll() Projects {
	var projects = Projects{}

	err := r.db.View(func(tx *bolt.Tx) error {

		// Assume bucket exists and has keys

		b := tx.Bucket(BucketName)
		b.ForEach(func(k, v []byte) error {
			var p Project
			err := json.Unmarshal(v, &p)
			if err != nil {
				log.Fatal(err)
			}
			projects = append(projects, p)
			return nil
		})

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return projects
}

// Add a Project to the Repository
func (r *ProjectRepo) Add(p Project) Project {

	err := r.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(BucketName)
		if err != nil {
			return err
		}

		p.ID, _ = b.NextSequence()
		project, err := json.Marshal(p)
		if err != nil {
			return err
		}

		// make a byte key from the ID
		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, p.ID)

		// save the project
		err = b.Put(key, project)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Saved Project with ID %d", p.ID)

	return p
}


// Save a Project to the Repository
func (r *ProjectRepo) Save(p Project) Project {

	err := r.db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(BucketName)
		if err != nil {
			return err
		}

		project, err := json.Marshal(p)
		if err != nil {
			return err
		}

		// make a byte key from the ID
		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, p.ID)

		// save the project
		err = b.Put(key, project)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return p
}

// Delete a Project from the Repository
func (r *ProjectRepo) Delete(id uint64) Project {
	var project Project
	err := r.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(BucketName)
		if bucket == nil {
			return fmt.Errorf("Projects Bucket %q not found!", BucketName)
		}

		// make a byte key from the ID
		key := make([]byte, 8)
		binary.BigEndian.PutUint64(key, id)

		// get the Project by its ID key from the DB
		json.Unmarshal(bucket.Get(key), &project)
		return bucket.Delete(key)
	})

	if err != nil {
		return Project{}
	}

	// return empty Projects list if not found
	return project
}
