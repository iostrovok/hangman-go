(function() {
    "use strict";

    angular.module('common')
        .service('GameDataService', GameDataService);

    GameDataService.$inject = ['$q', '$http', 'ApiPath'];

    function GameDataService($q, $http, ApiPath) {
        var service = this;

        service.checkedLogin = false;

        service.userData = {
            attempt: '00',
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
            service.userData.attempt = 0;
            service.userData.last_attempt = false;
            service.userData.last_letter = '';
            service.userData.last_word = '';
            service.userData.state = '';
            service.userData.view_word = '';
            service.userData.last_letters = '';
            service.userData.viewError = false;
        }


        service._setData = function(data) {
            console.log('SetData', data);
            if (data != undefined && data != null) {
                service.userData.attempt = Number(data.attempt);
                service.userData.last_attempt = data.last_attempt;
                service.userData.last_letter = data.last_letter;
                service.userData.last_word = data.last_word;
                service.userData.state = data.state;
                service.userData.view_word = data.view_word;
                service.userData.wins = data.wins;
                service.userData.losings = data.losings;
                service.userData.last_letters = data.last_letters;

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
                    case "bad_move":
                        service.userData.error_message = "No correct turn.";
                        service.userData.viewError = true;
                        break;
                    default:
                        service.userData.viewError = false;
                        service.userData.error_message = "";
                        break;
                }




                console.log('service.userData', service.userData);

            } else {
                service._cleanData();
            }
            service.userData.viewData = true;
        }

        service._send = function(url, method, params) {
            console.log("GameDataService._send()", url, method, params);
            $http({
                url: url,
                method: method,
                params: params,
            }).then(function(data) {
                /* if request is successful */
                service._setData(data.data);
            }, function(data) {
                /* if request is not successful */
                console.error("_send", url, method, params, data);
            });
        }

        service.Move = function(params) {
            console.log("-> GameDataService.NewGame()");
            service._send('/move', 'GET', params);
        }

        service.NewGame = function() {
            console.log("-> GameDataService.NewGame()");
            service._send('/start', 'GET', {});
        }

        service.Init = function() {
            console.log("-> GameDataService.Init()");
            service._send('/user_info', 'GET', {});
        };
    }



})();
