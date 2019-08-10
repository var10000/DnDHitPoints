<template>
  <f7-app id="app">
  <div class="room">
    <header class="header">
      <div class="header__text">
        <p class="header__text-title"> D&D craft page</p>
        <p class="header__text-title_description"> The  smartest app to organize your fight </p>
      </div>
      <div>
        <img
          v-on:click="isCreateCharacter = true"
          class="header__image"
          src="../assets/plus.png"/>
      </div>
    </header>

    <create-character-popup
            :opened="isCreateCharacter"
            @popup:completed="saveCharacter"
            @popup:closed="isCreateCharacter = false"/>
  </div>
  </f7-app>
</template>

<script>
  import 'bootstrap/dist/css/bootstrap.css'
  import 'bootstrap-vue/dist/bootstrap-vue.css'
  import CreateCharacterPopup from "./components/CreateCharacterPopup"

export default {
  name: 'app',
  components: {
    CreateCharacterPopup,
  },
  data() {
    return {
      isCreateCharacter: false,
    }
  },
  methods: {
    saveCharacter() {

    }
  }
}
</script>

<style>
  @import url('https://fonts.googleapis.com/css?family=Cinzel:700&display=swap');
  .room {
    background-image: url("../assets/background-dragon.jpg");
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center;
    text-align: center;
    width: 100%;
    min-height: 1080px;
    position: absolute;
    top: 0;
    right: 0;
    left: 0;
  }
  .header {
    display: flex;
    flex-direction: row;
    justify-content: space-between;
    color: white;
  }
  .header__text-title {
    font-family: 'Cinzel', serif;
    font-weight: lighter;
    letter-spacing: 2px;
    font-size: 40px;
    text-align: center;
  }
  .header__image {
    color: white;
    width: 40px;
  }
</style>
