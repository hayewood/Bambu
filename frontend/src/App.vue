<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <header>
      <nav>
        <Navbar></Navbar>
      </nav>
    </header>
    <main class="row s12 blue-grey darken-4">
      <Chat class="col s4" title="Channel 1"></Chat>
    </main>
    <Footer></Footer>
  </div>
</template>

<script src="http://code.jquery.com/jquery-1.11.0.min.js"></script>
<script src="https://cdnjs.cloudflare.com/ajax/libs/materialize/1.0.0/js/materialize.min.js"></script>

<script>
  import 'materialize-css'
  import Chat from './components/Chat.vue'
  import Footer from './components/Footer.vue'
  import Navbar from './components/Navbar.vue'
  import Vue from 'vue';
  import VeeValidate from 'vee-validate';
  import axios from 'axios';


  Vue.use(VeeValidate);

  export default {
    name: 'app',
    components: {
      Chat,
      Footer,
      Navbar
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
