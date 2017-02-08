(function() {
    "use strict";

    angular.module('common', [])
        .constant('ApiPath', 'http://localhost:80')
        .config(config);

    config.$inject = ['$httpProvider'];

    function config($httpProvider) {
        $httpProvider.interceptors.push('loadingHttpInterceptor');
    }

})();
