<template>
  <div class="row mb-4">
    <div class="col-12 card">
      <div class="card-body row justify-content-between">
        <div class="col-2 font-monospace align-self-center">
          <input class="form-control" v-model="path" placeholder="Path..." />
        </div>
        <div class="col align-self-center">
          <input class="form-control" v-model="url" placeholder="URL..." />
        </div>
        <div class="col-auto align-self-center">
          <button
            class="btn btn-outline-success mr-2"
            v-on:click="onSaveButtonClick"
          >
            Save
          </button>
          <button
            class="btn btn-outline-danger"
            v-on:click="onAbortButtonClick"
          >
            Abort
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      path: "",
      url: "",
    };
  },
  methods: {
    onSaveButtonClick: function() {
      this.$emit("startLoading");
      const formDataBody = new FormData();
      formDataBody.append("path", this.path);
      formDataBody.append("url", this.url);
      axios({
        method: "post",
        url: `http://${location.host}/api/shortcuts`,
        data: formDataBody,
        headers: {
          "Content-Type": "multipart/form-data",
          Authorization: localStorage.getItem("password"),
        },
      })
        .then((response) => {
          this.$emit("message", response.status);
          this.$emit("update");
          this.$emit("stopLoading");
          this.path = "";
          this.url = "";
        })
        .catch((error) => {
          this.$emit("message", error.response.status);
          this.$emit("update");
          this.editMode = true;
          this.$emit("stopLoading");
        });
    },
    onAbortButtonClick: function() {
      this.$emit("disableCreator");
    },
  },
};
</script>
