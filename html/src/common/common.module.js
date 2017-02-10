(function() {
    "use strict";

    angular.module('common', [])
        .config(config);
        /*.constant('ApiPath', 'http://localhost:80')*/

    config.$inject = ['$httpProvider'];

    function config($httpProvider) {
        $httpProvider.interceptors.push('loadingHttpInterceptor');
    }

})();
