(function() {
    "use strict";

    angular.module('public')
        .controller('GameController', GameController);

    GameController.$inject = ['$scope', '$stateParams', 'GameDataService'];

    function GameController($scope, $stateParams, GameDataService) {
        var $ctrl = this;
        $ctrl.letter = '';

        $ctrl.GameDataService = GameDataService;
        $ctrl.userData = GameDataService.userData;

        console.log('$stateParams: ', $stateParams);


        $ctrl.GameDataService.NewGame()


        $ctrl.submit = function() {
            var data = {
                'letter': $ctrl.letter
            };
            GameDataService.Move(data);
            $ctrl.letter = '';
        }

        $ctrl.newGame = function() {
            $ctrl.GameDataService.NewGame()
        }
    }

})();
