// Checks whether the web browser thinks that a user is logged in and
// sets the Login/Logout button accordingly
var login = new function() {
    $(document).ready(function() {
        $.ajax({
            type:  "GET",
            url:   "api/v1/user",
            cache: false,
            success: function(user){
                try {
                    console.log(user);
                    if(! user){
                        location.href = "/restricted";
                    }else{
                        localStorage.setItem("name",      user.name);
                        localStorage.setItem("username",  user.username);
                        localStorage.setItem("email",     user.email);
                    }
                    login.setLoginButton();
                }
                catch(err){
                    console.log("Error while parsing:");
                    console.log(err);
                    location.href = "/restricted";
                }
            },
            error: function(request, status, error){
                if(request.status == "404"){
                    /*$("#menu").after(
                        '<div class="alert alert-warning alert-dismissable">'+
                            '<button type="button" class="close" ' + 
                            'data-dismiss="alert" aria-hidden="true">' + 
                            '&times;' + 
                            '</button>' + 
                            'Cannot connect to API server. We cannot check whether the user is logged in or not. ' +
                        '</div>'
                    );*/
                    login.setLoginButton();
                }else if (request.status == "401"){
                    console.log("No user logged in, clear Local Storage for this App");
                    localStorage.removeItem("name");
                    localStorage.removeItem("username");
                    localStorage.removeItem("email");
                    login.setLoginButton();
                }
            }
        });
    });
    
    // sets the login button to either Login or Logout
    this.setLoginButton = function() {
        var uri = window.location.pathname;
        if(uri.match(/restricted/)){
            var name = localStorage.name;
            $("#login-button").html('<li><a href="/restricted">Logged in as '+name+'</a></li><li><a href="/officesso-login/logout">Logout</a></li>');
            return;
        }else{
            $('#login-button').html('<li><a href="/restricted">Not logged in</a></li><li><a href="/restricted">Login</a></li>');
        }
    }
}