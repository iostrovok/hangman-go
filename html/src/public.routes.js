(function() {
    'use strict';

    angular.module('public')
        .config(routeConfig);

    /**
     * Configures the routes and views
     */
    routeConfig.$inject = ['$stateProvider'];

    function routeConfig($stateProvider) {
        // Routes
        $stateProvider
            .state('public', {
                absract: true,
                templateUrl: '/html/src/public.html'
            })
            .state('public.home', {
                url: '/',
                templateUrl: '/html/src/home/home.html',
                controller: 'HomeController',
                controllerAs: 'homeCtrl'
            })
            .state('public.game', {
                url: '/game/',
                templateUrl: '/html/src/game/game.html',
                controller: 'GameController',
                controllerAs: 'gameCtrl'
            });
    }
})();
