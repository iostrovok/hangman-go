(function() {
    "use strict";

    angular.module('public').controller('GameController', GameController);

    GameController.$inject = ['GameDataService'];

    function GameController(GameDataService) {
        var $ctrl = this;
        $ctrl.letter = '';
        $ctrl.abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ".split('');

        $ctrl.GameDataService = GameDataService;
        $ctrl.userData = GameDataService.userData;

        $ctrl.abc_list = [];
        var $letter_post_func = function() {
            var maxId = parseInt(1 + $ctrl.abc.length / 4) * 4;
            for (var i = 0; i < maxId; i++) {

                var row = parseInt(i / 4);
                var col = parseInt(i % 4);

                if ($ctrl.abc_list[row] == null) {
                    $ctrl.abc_list[row] = [];
                }

                if ($ctrl.abc[i] != null) {
                    $ctrl.abc_list[row][col] = {
                        val: $ctrl.abc[i],
                        show: $ctrl.view_last_letter($ctrl.abc[i])
                    }
                }
            }

        };

        $ctrl.Move = function(letter) {
            if (!$ctrl.userData.inGame) {
                return
            }

            var data = {
                'letter': letter
            };
            GameDataService.Move(data, $letter_post_func);
            $ctrl.letter = '';
        }

        $ctrl.newGame = function() {
            $ctrl.GameDataService.NewGame($letter_post_func);
        }

        $ctrl.view_last_letter = function(letter) {
            letter = String(letter).toUpperCase();
            for (var i in $ctrl.userData.last_letters) {
                if (String($ctrl.userData.last_letters[i]).toUpperCase() == letter) {
                    return false;
                }
            }
            return true;
        }

        $ctrl.GameDataService.NewGame($letter_post_func);
    }
})();
