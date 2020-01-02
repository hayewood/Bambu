<template>
    <!-- Modal Structure -->
    <div id="sign-in-modal" class="modal profile-overlay blue-grey darken-4">
        <div class="modal-header">
            <h1 class="brand-logo center" style="position:relative"> Login to Bambu </h1>   
        </div> 
            <div class="modal-content blue-grey darken-4"> 
                <div class="col">
                    <div class = "col s12">
                        <ul id="login-tabs" class="tabs tabs-fixed-width tab-demo blue-grey darken-4 z-depth-1">
                            <li class="active tab col s6"><a class="white-text" href="#login-tab">Login</a></li>
                            <li class="tab col s6"><a class="white-text" href="#signin-tab">Sign Up</a></li>
                        </ul>  
                    </div>

                    <div id="login-tab" class="col s12">

                        <div class="input-field col s4 box"  style="margin-top:15px;margin-bottom:5px;">
                            <input type="email" v-model.trim="mEmail" placeholder="Email">
                        </div>
                        <div class="input-field col s4 box">
                            <input type="text" v-model.trim="mUsername" placeholder="Username">
                        </div>
                        <a class="btn lime lighten-2 col black-text center-align" style="line-height: 2rem;" @click="sign_in"> Log In </a>

                        <g-signin-button class="col lime lighten-2 black-text"
                            :params="googleSignInParams"
                            @success="onSignInSuccess"
                            @error="onSignInError">
                            <img src="../assets/icons8-google.svg"> Sign in with google
                        </g-signin-button>
                    </div>


                    <div id="signin-tab" class="col s12">
                        <div class="input-field col s4 box" style="margin-top:15px;margin-bottom:0px;">
                            <input type="text" v-model.trim="mUsername" placeholder="Username">
                        </div>
                        <div class="input-field col s4 box" style="margin-top:5px;margin-bottom:5px">
                            <input type="password" placeholder="Password">
                        </div>
                        <div class="input-field col s4 box" style="margin-bottom:5px">
                            <input type="email" v-model.trim="mEmail" placeholder="Email">
                        </div>
                        <vue-recaptcha class="col" sitekey='6LdLjssUAAAAAMscS5yyO5z7k6QAxxW5DI9dn0PB'></vue-recaptcha>
                        <a class="btn lime lighten-2 col black-text center-align" style="line-height: 2rem;margin:5px;" @click="sign_in"> Log In </a>

                    </div>
                </div>
            </div>
    </div>
</template>

<script src="https://apis.google.com/js/platform.js" async defer></script>

<script>
  import VueRecaptcha from 'vue-recaptcha';
  import $ from 'jquery'

export default {
  name: 'ProfileOverlay',
    components: {
        VueRecaptcha
    },
    props: {
        signIn: Boolean,
        username: String,
        profilePic: String,
        email: String
    },
  data: function() {
    return {
        mEmail: null, // Email address used for grabbing an avatar
        mUsername: null, // Our username
        mProfilePic: null,
        logged_in: this.signIn,
        loginOverlay: null,
        tabInstance: null,
        googleSignInParams: {
        client_id: '993709220930-f8p9d6m171va04llfk66l7ujj7kq7fuj.apps.googleusercontent.com',
        modal_header: "Login"
      }
    }
  },

  mounted() {
        var tabs = document.querySelectorAll('.tabs')
        this.tabInstance = tabs;
        for (var i = 0; i < tabs.length; i++){
            M.Tabs.init(tabs[i]);
        }
        var modals = document.querySelectorAll('.modal');
        this.loginOverlay = modals;
        for (var i = 0; i < modals.length; i++){
            M.Modal.init(modals[i]);
            
        }
  },

    created: function() {
    },

    methods: {
    sign_in: function() {
            if (!this.mEmail) {
                Materialize.toast('You must enter an email', 2000);
                return
            }
            if (!this.mUsername) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            this.mEmail = $('<p>').html(this.mEmail).text();
            this.mUsername = $('<p>').html(this.mUsername).text();
            this.mProfilePic = $('<p>').html(this.mProfilePic).text();
            console.log("signin"+ this.logged_in);
            this.logged_in = true;
            this.open();
        },
     onSignInSuccess (googleUser) {
        // `googleUser` is the GoogleUser object that represents the just-signed-in user.
        // See https://developers.google.com/identity/sign-in/web/reference#users
        const profile = googleUser.getBasicProfile() // etc etc
        console.log('ID: ' + profile.getId()); // Do not send to your backend! Use an ID token instead.
        console.log('Name: ' + profile.getName());
        console.log('Image URL: ' + profile.getImageUrl());
        console.log('Email: ' + profile.getEmail());
        this.mEmail = profile.getEmail();
        this.mUsername = profile.getName();
        this.mProfilePic = profile.getImageUrl();
        this.logged_in = true;
        this.open();

    },
    onSignInError (error) {
      // `error` contains any error occurred.
      console.log('OH NOES', error)
    },

    sign_out: function() {
        var self = this;
        var auth2 = gapi.auth2.getAuthInstance();
        auth2.signOut().then(function () {
            console.log('User signed out.');
            self.mEmail = null;
            self.mUsername = null;
            self.mProfilePic = null;
            self.logged_in = false; 
        });

    },

    open: function() {
        var self = this;

        var modalInstance = M.Modal.getInstance(this.loginOverlay[0]);
        modalInstance.close();
        this.$emit('login',this.logged_in);
        this.$emit('login_info',[this.mEmail,this.mUsername, this.mProfilePic]);
        //for(var i = 0; i < this.loginOverlay.length; ++i) {
         //   this.loginOverlay[i].close();
        //}
    }
    }
}
</script>

<style>
    .g-signin-button {
    /* This is where you control how the button looks. Be creative! */
    padding: 4px 8px;
    border-radius: 3px;
    background-color: #4285F4;
    color: #000000;
    box-shadow: 0 3px 0 #4285F4;
        }

    .modal.profile-overlay {
        width: fit-content;
        height: fit-content;
    }

    .tabs .indicator { display: none; }
    .tabs .tab a.active { border-bottom: 2px solid #f6b2b5; }

</style> 