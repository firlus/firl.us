<template>
    <div class="row mb-4">
        <div class="col-12 card">
            <div class="card-body row justify-content-between">
                <div class="col-2 font-monospace align-self-center">{{ path }} </div>
                <div class="col align-self-center font-monospace" v-if="!editMode">
                    {{ shortenUrl(url) }}
                </div>
                <div class="col align-self-center" v-else>
                    <input class="form-control" v-model="changedUrl">
                </div>
                <div class="col-auto align-self-center" v-if="!editMode">
                    <button class="btn btn-outline-primary" v-on:click="onEditButtonClick">Edit</button>
                </div>
                <div class="col-auto" v-else>
                    <button class="btn btn-outline-success mr-2" v-on:click="onSaveButtonClick">Save</button>
                    <button class="btn btn-outline-warning mr-3" v-on:click="onAbortButtonClick">Abort</button>
                    <button class="btn btn-outline-danger" v-on:click="onDeleteButtonClick">Delete</button>
                </div>
            </div>
        </div>
    </div>
</template>


<script>
    import axios from 'axios';

    export default {
        props: ['path', 'url'],
        data() {
            return {
                editMode: false,
                changedUrl: ''
            }
        },
        methods: {
            onEditButtonClick: function() {
                this.changedUrl = this.url;
                this.editMode = true;
            },
            onSaveButtonClick: function() {
                this.$emit('startLoading');
                const formDataBody = new FormData();
                formDataBody.append('url', this.changedUrl);
                axios({
                    method: 'put',
                    url: `https://firl.us/api/shortcuts/${this.path}`,
                    data: formDataBody,
                    headers: {'Content-Type': 'multipart/form-data'}
                }).then((response) => {
                    this.$emit('message', response.status);
                    this.$emit('update');
                    this.editMode = false;
                    this.$emit('stopLoading');
                }).catch((error) => {
                    this.$emit('message', error.response.status);
                    this.$emit('update');
                    this.editMode = true;
                    this.$emit('stopLoading');
                });
            },
            onDeleteButtonClick: function() {
                this.$emit('startLoading');
                axios({
                    method: 'delete',
                    url: `https://firl.us/api/shortcuts/${this.path}`,
                    headers: {'Content-Type': 'multipart/form-data'}
                }).then((response) => {
                    this.$emit('message', response.status);
                    this.$emit('update');
                    this.editMode = false;
                    this.$emit('stopLoading');
                }).catch((error) => {
                    this.$emit('message', error.response.status);
                    this.$emit('update');
                    this.editMode = true;
                    this.$emit('stopLoading');
                });
            },
            onAbortButtonClick: function() {
                this.editMode = false;
            },
            shortenUrl: (url) => {
                if(url.length < 75) {
                    return url;
                } else {
                    return url.substring(0, 74).concat('...');
                }
            }
        }
    }
</script>