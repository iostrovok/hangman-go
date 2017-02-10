(function() {
    "use strict";

    angular.module('public').controller('HomeController', HomeController);

    HomeController.$inject = ['$scope', 'GameDataService'];

    function HomeController($scope, GameDataService) {
        var $ctrl = this;
        $ctrl.isPlayNow = false;
        $ctrl.GameDataService = GameDataService;
        $ctrl.userData = GameDataService.userData;

        var init = function() {
            $ctrl.isPlayNow = $ctrl.userData.state == 'is_play_now' ? true : false;
        }

        $ctrl.GameDataService.Init(init);
    }

})();
