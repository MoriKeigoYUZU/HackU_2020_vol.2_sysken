<template>
  <div class="home">
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <v-container>
      <Title titlename="スッカ" size="6" />
      <v-row>
      <v-col offset="5">
      <HomeIcon name="ToDoリスト" color="#F29993" pash="/Todo" />
      </v-col>
      <v-col>
      <HomeIcon name="カレンダー" color="#97D5C5" pash="/Calendar" />
      </v-col>
      <v-col offset="5">
      <HomeIcon name="育成" color="#f8b739" pash="/Raising" />
      </v-col>
      </v-row>
    </v-container>
  </div>
</template>


<script>
// @ is an alias to /src
import HomeIcon from "@/components/HomeIcon.vue";
import Title from "@/components/Title.vue";

export default {
  name: "Home",
  components: {
    HomeIcon,
    Title,
  },
};
</script>
<style>
.theme--light.v-divider {
  border-color: #6b421d !important;
  border-width: 5px !important;
}
h1 {
  color: #6b421d !important;
}
a{
  text-decoration: none;
}
</style>
