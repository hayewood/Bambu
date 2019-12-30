<template>
    <!-- Modal Structure -->
    <div id="profileModal" class="modal">
        <div class="modal-content">                    
            <div class="row" v-if="!signed_in">
                <h1 class="black-text center">Login</h1>
                <div class="input-field col s4">
                    <input type="email" v-model.trim="email" placeholder="Email">
                </div>
                <div class="input-field col s4">
                    <input type="text" v-model.trim="username" placeholder="Username">
                </div>
                <div class="input-field col s4">
                    <button class="waves-effect lime lighten-2 btn" @click="sign_in">
                        <i class="material-icons right" style="line-height: 40px;" >done</i>Join
                    </button>
                </div>
            </div>
            <div class="row" v-if="signed_in">
                    <h1 class="black-text center">Welcome!</h1>
                    <div class="modal-footer">
                        <a href="#!" class="modal-close waves-effect waves-green btn-flat">Close</a>
                    </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
  name: 'ProfileOverlay',

  data: function() {
    return {
        email: null, // Email address used for grabbing an avatar
        username: null, // Our username
        signed_in: false, // Are we signed in
        loginOverlay: null,
    }
  },

    created: function() {
        document.addEventListener("DOMContentLoaded", function() {
            var modals = document.querySelectorAll('.modal');
            M.Modal.init(modals);
        });
    },

    methods: {
    sign_in: function() {
            if (!this.email) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.email = $('<p>').html(this.email).text();
            this.username = $('<p>').html(this.username).text();
            this.signed_in = true;
            this.open();
        },

            open: function() {
                var self = this;
                if (location.hostname === "localhost" || location.hostname === "127.0.0.1")
                    this.ws = new WebSocket(((window.location.protocol === "https:") ? "wss://" : "ws://") + "localhost:8000" + "/ws");
                else
                    this.ws = new WebSocket(((window.location.protocol === "https:") ? "wss://" : "ws://") + "calculator-example.herokuapp.com" + "/ws");

                this.ws.addEventListener('message', function(event) {
                    var msg = JSON.parse(event.data);
                    self.chatContent += '<div class="chip lime lighten-2">'
                            + '<i class="material-icons left black-text">account_circle</i>'
                                + msg.username
                            + '</div>'
                        + emojione.toImage(msg.message) + '<br/>'; // Parse emojis
                    var element = document.getElementById('chat-messages');
                    element.style.color = "#FFFFFF";
                    element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
                });

                this.ws.addEventListener('close', (event) => {
                    this.maybeReconnectToWebsocket(event);
                });
            }
    }
}
</script>