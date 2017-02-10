(function() {
    "use strict";

    angular.module('common')
        .service('GameDataService', GameDataService);

    GameDataService.$inject = ['$http'];

    function GameDataService($http) {
        var service = this;

        service.checkedLogin = false;

        service.userData = {
            attempt: '0',
            viewError: "",
            last_attempt: false,
            last_letter: "",
            last_word: "",
            losings: 0,
            state: "",
            view_word: "",
            wins: 0,
            inGame: false,
            isWinner: false,
            isLoser: false,
            last_letters: ''
        };

        service._cleanData = function(data) {
            service.userData.attempt = '0';
            service.userData.last_attempt = false;
            service.userData.last_letter = '';
            service.userData.last_word = '';
            service.userData.state = '';
            service.userData.view_word = '';
            service.userData.last_letters = '';
            service.userData.viewError = false;
        }

        service._setData = function(data) {

            if (data != undefined && data != null) {

                for (var t in data ) {
                    service.userData[t] = data[t];
                }

                service.userData.inGame = false;
                service.userData.isWinner = false;
                service.userData.isLoser = false;
                switch (data.state) {
                    case "is_play_now":
                        service.userData.inGame = true;
                        break;
                    case "winner":
                        service.userData.isWinner = true;
                        break;
                    case "loser":
                        service.userData.isLoser = true;
                        break;
                }

                switch (data.error) {
                    case "duplicate":
                        service.userData.error_message = "You've already used this letter.";
                        service.userData.viewError = true;
                        break;
                    case "bad_move":
                        service.userData.error_message = "No correct turn.";
                        service.userData.viewError = true;
                        break;
                    case "not_in_the_game":
                        service.userData.error_message = "No correct turn.";
                        service.userData.viewError = true;
                        break;
                    default:
                        service.userData.viewError = false;
                        service.userData.error_message = "";
                        break;
                }

            } else {
                service._cleanData();
            }
            service.userData.viewData = true;
            // console.log("service.userData: ", service.userData);
        }

        service._send = function(url, method, params, post_load) {

              $http({
                url: url,
                method: method,
                params: params,
            }).then(function(data) {
                /* if request is successful */
                service._setData(data.data);
                if (post_load != null ) {
                    post_load()
                }
            }, function(data) {
                /* if request is not successful */
            });
        }

        service.Move = function(params, post_load) {
            service._send('/move', 'GET', params, post_load);
        }

        service.NewGame = function(post_load) {
            service._send('/start', 'GET', {}, post_load);
        }

        service.Init = function(post_load) {
            service._send('/user_info', 'GET', {}, post_load);
        };
    }
})();
