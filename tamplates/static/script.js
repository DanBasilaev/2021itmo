var app = new Vue({
    el: '#app',
    data: {
        header: 'Защити свои данные',
        sign_up: 'Регистрация',
        log_in: 'Вход',

        form: '',
        encoding: 'Кодировать',

        name: '',
        lastname: '',
        login: '',
        password: '',
        registration: 'Регистрация',
        log: 'Войти',
        exit: 'Выход'
    },
    methods:{
        regBtn: function () {
            var regbtn = document.getElementById('reg').hidden = false
            var user= document.getElementById('user').hidden = true
            var logbtn = document.getElementById('log').hidden = true

        },
        logBtn: function () {
            var logbtn = document.getElementById('log').hidden = false
            var user= document.getElementById('user').hidden = true
            var regbtn = document.getElementById('reg').hidden = true

        },
        cleanMessage: function () {
            try {
                if (this.form === ''){
                    throw 'error'
                }
                else {
                    this.form = '';
                }
            }
            catch{
                alert("")
            }
        },
        regBtn_menu: function () {
            var regbtn = document.getElementById('reg').hidden = true
            var user= document.getElementById('user').hidden = false
            var sign = document.getElementById("sign-log").hidden = true
            var user_name = document.getElementById('name').hidden = false

        },
        logBtn_menu: function () {
            var logbtn = document.getElementById('log').hidden = true
            var user= document.getElementById('user').hidden = false
            var sign = document.getElementById("sign-log").hidden = true
            var user_name = document.getElementById('name').hidden = false
        },

        exitBtn: function (){
            var user= document.getElementById('user').hidden = true
            var sign = document.getElementById("sign-log").hidden = false
            var user_name = document.getElementById('name').hidden = true

        }
    }
})

$(document).keypress(
    function(event){
        if (event.which == '13') {
            event.preventDefault();
        }
    });
