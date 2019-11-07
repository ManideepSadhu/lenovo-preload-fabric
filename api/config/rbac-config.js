'use strict';
module.exports = {
    "admin": {
        "can": [
            "GET:/softwarelicense/*", 
            "GET:/softwarelicenselist", 
            "GET:/softwarelicensehistory/*", 
            "POST:/softwarelicense", 
            "PUT:/softwarelicense", 
            "DELETE:/softwarelicense/*"
        ]
    }
}