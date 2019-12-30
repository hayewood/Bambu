<template>
    <div class="nav-wrapper blue-grey darken-4 z-depth-5">
            <a href="/" class="brand-logo center hoverable" style="font-family:'Monoton',serif; font-size: 49px; color: #dce775;">BAMBU</a>
            
            <div id="profileDropdown">
                 <!-- Dropdown Trigger -->
                <a class="dropdown-trigger right btn-floating lime lighten-2 hoverable" id="dropdown-button" style="margin: 10px;" data-target="dropdown1">
                    <i class="material-icons" style="line-height: 40px; color: black;">account_circle</i>
                </a>
                <!-- Dropdown Structure -->
                <ul id='dropdown1' class='dropdown-content lime lighten-2' style="z-index: 7;"> 
                    <li>
                        <a href="#profileModal" class="modal-trigger black-text modal-close"><i class="material-icons">account_circle</i>Account</a>    
                    </li>
                 </ul>
            </div>
            <ProfileOverlay></ProfileOverlay>
        </div>
</template>

<script>
    import ProfileOverlay from './ProfileOverlay.vue'
    export default {
    name: 'Navbar',

    components: {
        ProfileOverlay
    },

    data: function() {
        return {
            email: null, // Email address used for grabbing an avatar
            username: null, // Our username
            signed_in: false, // Are we signed in
        }
    },

        created: function() {
            document.addEventListener("DOMContentLoaded", function() {
                var dropdownTriggers = document.querySelectorAll('.dropdown-trigger');
                M.Dropdown.init(dropdownTriggers,{ constrainWidth: false, belowOrigin: true});
            });
        }
    }

</script>

<style>
  @import "https://fonts.googleapis.com/icon?family=Material+Icons";
</style>
