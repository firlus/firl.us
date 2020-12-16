<template>
  <div>
    <div class="row mt-5 mb-4">
      <div class="col-12">Login</div>
    </div>
    <div class="row justify-content-center">
      <div class="col">
        <input type="password" class="form-control" v-model="password" />
        <button class="btn btn-outline-primary" @click="login">Login</button>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      password: "",
    };
  },
  methods: {
    login: function() {
      axios
        .get(`http://${location.host}/api/shortcuts`, {
          headers: {
            Authorization: this.password,
          },
        })
        .then((_) => {
          localStorage.setItem("password", this.password);
          this.$router.push("/list");
        })
        .catch((error) => console.log(error));
    },
  },
};
</script>

<style></style>
