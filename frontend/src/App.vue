<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <header>
      <nav>
        <Navbar :signIn="signed_in" :email="email" :profilePic="profilePic" :username="username" v-on:signed_in="signedInCallback($event)" v-on:login_info="loginInfoCallback($event)"></Navbar>
      </nav>
    </header>
    <main class="row s12 blue-grey darken-4">
      <draggable handle=".handle">
        <Chat :signIn="signed_in" :email="email" :profilePic="profilePic" :username="username" v-on:signed_in="signedInCallback($event)" v-on:login_info="loginInfoCallback($event)" class="col s4 blue-grey darken-4" title="Channel 1"></Chat>
        <Chat :signIn="signed_in" :email="email" :profilePic="profilePic" :username="username" v-on:signed_in="signedInCallback($event)" v-on:login_info="loginInfoCallback($event)" class="col s4 blue-grey darken-4" title="Channel 2"></Chat>
        <Chat :signIn="signed_in" :email="email" :profilePic="profilePic" :username="username" v-on:signed_in="signedInCallback($event)" v-on:login_info="loginInfoCallback($event)" class="col s4 blue-grey darken-4" title="Channel 3"></Chat>
      </draggable>

    </main>
    <Footer></Footer>
  </div>
</template>

<script src="https://cdnjs.cloudflare.com/ajax/libs/jquery/1.12.4/jquery.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>

<script>
  import 'materialize-css'
  import DataStore from './DataStore'
  import Sortable from 'sortablejs';
  import draggable from 'vuedraggable'
  import Chat from './components/Chat.vue'
  import Footer from './components/Footer.vue'
  import Navbar from './components/Navbar.vue'
  import Vue from 'vue';
  
  Vue.use(draggable)

  export default {
      name: 'app',
      components: {
        Chat,
        draggable,
        Footer,
        Navbar
      },
        data:function() {
          return {
            isHoverDragable: false,
            title: null,
            signed_in: false,
            username: null,
            profilePic: null,
            email: null

          }
      },
      methods: {
        signedInCallback: function() {
          this.signed_in = true;
        },

        loginInfoCallback: function(event) {
            this.email = event[0];
            this.username = event[1];
            this.profilePic = event[2];

        }
      }
  }
</script>

<style>
  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    color: #2c3e50;
  }
  @import "https://cdnjs.cloudflare.com/ajax/libs/emojione/2.2.7/assets/css/emojione.min.css";
  @import "https://fonts.googleapis.com/icon?family=Material+Icons";
  @import "https://cdn.jsdelivr.net/emojione/2.2.6/assets/css/emojione.min.css";
  @import "https://fonts.googleapis.com/css?family=Monoton";
  @import "https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/css/materialize.min.css";

</style>
