<template>
 <div id="draggable-chat" style="border: 1px solid #424242;">
    <div class="col">     
        <div class="row-content center z-depth-1" style="border-bottom: 1px solid #424242">
             <span class="flow-text white-text center">{{ title }}</span>
            <i id="dragable-button" class="material-icons right lime-text lighten-2 handle">drag_indicator</i>
        </div>

        <div id="main-chat" class="row" style="position: relative; margin-bottom: 0px;">
            <div id="chat-messages" class="card-content left" style="margin: 10px;" v-html="chatContent">
            </div>
        </div>
    </div>

    <div id="input-box" class="col blue-grey darken-4 z-depth-1">
            <!-- if a session is joined -->
        <div class="row chat-box" style="margin-bottom:0px;margin-left:5px" v-if="signIn">
            <div class="input-field col s4 left box">
                <input type="text" v-model="newMsg" placeholder="Say something here" @keyup.enter="send">
            </div>
            <a class="btn-floating lime lighten-2" style="margin: 10px; margin-top:20px" @click="send">
                <i class="material-icons" style="line-height: 40px; color: black;">send</i>
            </a>
            <a class="btn-floating lime lighten-2" style="margin: 10px; margin-top:20px" @click="close">
                <i class="material-icons" style="line-height: 40px; color: black;">close</i>
            </a>
        </div>

        <!-- if a session is not joined -->
        <div class="row chat-box" style="margin-bottom:0px; margin-left:5px" v-if="!signIn">
            <div class="input-field col s4 box">
                <input type="email" v-model.trim="mEmail" placeholder="Email">
            </div>
            <div class="input-field col s4 box" style="margin-left:5px">
                <input type="text" v-model.trim="mUsername" placeholder="Username">
            </div>
          <a class="btn-floating lime lighten-2" style="margin: 10px; margin-top:20px" @click="join">
                <i class="material-icons" style="line-height: 40px; color: black;">done</i>
            </a>
        </div>
    </div>
</div>

</template>

<script>
import * as emojione from "emojione";
import $ from "jquery";
import Vue from 'vue';

export default {
  name: 'Chat',

  props: {
    title: String,
    username: String,
    profilePic: String,
    email: String,
    signIn: Boolean
  },

    watch: {
        signIn: function () {
            this.open();
        },
    },
  data: function() {
    return {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        mEmail: null, // Email address used for grabbing an avatar
        mUsername: null, // Our username
        mProfilePic: null,
        joined: null, // True if email and username have been filled in
        autoReconnectInterval: 100,
    }
  },

  created: function(){
    var self = this;
    // Created another Vue instance
    window.Event = new Vue();

    document.addEventListener("DOMContentLoaded", function() {
        // Changing the cursor for the chat 
        var drag_chat = document.querySelectorAll("[id='dragable-button']");
        for(var i = 0; i < drag_chat.length; i++) {
            drag_chat[i].style.cursor='all-scroll';
        }

        // Changing the cursor for the input box
        var input_box = document.querySelectorAll("[id='input-box']");
        for(var i = 0; i < input_box.length; i++) 
            input_box[i].style.cursor='default';

        
    });
  },

    methods: {
        send: function () {
           if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        email: this.email,
                        picture: this.profilePic,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                this.newMsg = ''; // Reset newMsg
           }
        },

        join: function () {
            if (!this.mEmail) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.mUsername) {
                Materialize.toast('You must choose a username', 2000);
                return
            }

            this.mMail = $('<p>').html(this.mEmail).text();
            this.mProfilePic = $('<p>').html(this.mProfilePic).text();
            this.mUsername = $('<p>').html(this.mUsername).text();
            this.joined = true;
            this.$emit('signed_in',this.joined);
            this.$emit('login_info',[this.mEmail,this.mUsername, this.mProfilePic]);
            //this.open();
        },

        profileURL: function() {
            return '../assets/logo.png';
        },

        maybeReconnectToWebsocket: function(event) {
            switch (event.code){
                case 1000:	// CLOSE_NORMAL
                    console.log("WebSocket: closed");
                    break;
                default:	// Abnormal closure
                    this.reconnect(event);
                    break;
                }
                //this.onclose(e);
        },

        reconnect: function(event) {
            console.log(`WebSocketClient: retry in ${this.autoReconnectInterval}ms`,event);
            this.ws.addEventListener('message',null);
            this.ws.addEventListener('close',null);
            this.ws = null;
            var that = this;
            setTimeout(function(){
                console.log("WebSocketClient: reconnecting...");
                that.open();
            },this.autoReconnectInterval);
        },

        open: function() {
            var self = this;
            if (location.hostname === "localhost" || location.hostname === "127.0.0.1")
                this.ws = new WebSocket(((window.location.protocol === "https:") ? "wss://" : "ws://") + "localhost:8000" + "/ws");
            else
                this.ws = new WebSocket(((window.location.protocol === "https:") ? "wss://" : "ws://") + "calculator-example.herokuapp.com" + "/ws");

            this.ws.addEventListener('message', function(event) {
                var msg = JSON.parse(event.data);

                if(msg.picture != "") {
                    self.chatContent += '<div class="chip lime lighten-2">'
                            + '<img src="' + msg.picture + '">' // Avatar
                            + msg.username
                            + '</div>'
                        + emojione.toImage(msg.message) + '<br/>'; // Parse emojis
                    var element = document.getElementById('chat-messages');
                    element.style.color = "#FFFFFF";
                    element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
                }
                else {
                    self.chatContent += '<div class="chip lime lighten-2">'
                            + '<i class="material-icons left black-text">account_circle</i>'
                                + msg.username
                            + '</div>'
                        + emojione.toImage(msg.message) + '<br/>'; // Parse emojis
                    var element = document.getElementById('chat-messages');
                    element.style.color = "#FFFFFF";
                    element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
                }
            });

            this.ws.addEventListener('close', (event) => {
                this.maybeReconnectToWebsocket(event);
            });
            },

        close: function() {
            this.ws.close(1000);
            this.joined = false;
        }
    }
}
</script>

<style >


#chat-messages {
    min-height: 10vh;
    height: 60vh;
    width: 100%;
    overflow-y: scroll;
    color:white;
}

.box {
    padding:5px;
    border: 1px solid #616161;
     border-radius: 4px;
     background-color:#37474f;
}

input, select, textarea{
    color: #FFFFFF;
}

.material-icons.left {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    vertical-align: middle;
    line-height: inherit;
}

.material-icons.right {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    vertical-align: middle;
    line-height: inherit;
}
.row.chat-box {
    margin-right: 0px;
    margin-left: 0px;
}

.row.col {
    padding: 0px;
}

.chat-app {
    flex: 0 1 auto;
}
</style>
