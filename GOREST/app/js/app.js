(function(){
    var app = angular.module('adminp', [ ]);

    app.controller('adminpController', [ '$http', function($http){
        var userlist = this;

        userlist.users = [];

        $http.get('http://localhost:3000/json/').success(function(data){
            userlist.users = data;
        });
    }]);
})();
