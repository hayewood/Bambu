<template>
    <div class="nav-wrapper blue-grey darken-4 z-depth-5">
            <a href="/" class="brand-logo center hoverable" style="font-family:'Monoton',serif; font-size: 49px; color: #dce775;">BAMBU</a>
            
            <div id="profileDropdown">
                 <!-- Dropdown Trigger -->
                <a class="dropdown-trigger right btn-floating lime lighten-2 hoverable" id="dropdown-button" style="margin: 10px;" data-target="dropdown1">
                    <i class="material-icons" v-if="profilePic == null || profilePic == ''" style="line-height: 40px; color: black;">account_circle</i>
                    <img v-else-if="profilePic !=null || profilePic != ''" v-bind:src='profilePic' style='height: 100%; width: 100%; object-fit: contain'> 
                </a>
                <!-- Dropdown Structure -->
                <ul id='dropdown1' class='dropdown-content lime lighten-2 z-depth-5' style="z-index: 7;"> 
                    <li>
                        <a href="#sign-in-modal" class="modal-trigger black-text modal-close"><i class="material-icons">account_circle</i>Sign In</a>    
                    </li>
                 </ul>
            </div>
            <ProfileOverlay v-on:login="signInCallback($event)"  v-on:login_info="loginInfoCallback($event)" v-bind:signIn="false"></ProfileOverlay>
        </div>
</template>

<script>
    import ProfileOverlay from './ProfileOverlay.vue';
    import GSignInButton from 'vue-google-signin-button';
    import Vue from 'vue';
    Vue.use(GSignInButton)

    export default {
    name: 'Navbar',
    props: {
        signIn: {
            type: Boolean
        },
        username: {
            type: String
        },
        profilePic: {
            type: String
        },
        email: {
            type:String
        }
    },
    watch: {
        signIn: function () {
            this.signIn = this.signIn;
        }
    },
    components: {
        ProfileOverlay
    },

    data: function() {
        return {
            logged_in: this.signIn
        }
    },

        created: function() {
            document.addEventListener("DOMContentLoaded", function() {
                var dropdownTriggers = document.querySelectorAll('.dropdown-trigger');
                M.Dropdown.init(dropdownTriggers,{ constrainWidth: false, coverTrigger: false});
            });
        },
        methods: {
            signInCallback: function(event) {
                this.logged_in = true;
                this.$emit('signed_in',this.logged_in);
            },

            loginInfoCallback: function(event) {
                this.$emit('login_info',event);
            }
        }
    }

</script>

<style>
  @import "https://fonts.googleapis.com/icon?family=Material+Icons";
</style>
