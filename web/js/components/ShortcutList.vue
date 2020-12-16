<template>
    <div>
		<div class="row mt-5 mb-4 justify-content-between">
            <div class="col-auto">
                <h1>firl.us</h1>
            </div>
            <div v-if="loading" class="col-auto align-self-center">
                <div class="spinner-border" role="status">
                </div>
            </div>
		</div>
        <shortcut
            v-for="shortcut in list"
            :key="shortcut.Path"
            :path=shortcut.Path
            :url=shortcut.URL
            @update="updateList"
            @startLoading="startLoading"
            @stopLoading="stopLoading"
            @message="displayMessage">
        </shortcut>
        <div class="row" v-if="!creator">
            <div class="col">
                <button v-on:click="enableCreator" class="btn btn-outline-primary">
                    New shortcut
                </button>
            </div>
        </div>
        <shortcut-creator v-else @disableCreator="disableCreator"
            @update="updateList"
            @startLoading="startLoading"
            @stopLoading="stopLoading"
            @message="displayMessage">
        </shortcut-creator>
	</div>
</template>

<script>
    import axios from 'axios';

    export default {
        data() {
            return {
                list: null,
                loading: false,
                creator: false,
            }    
        },
        mounted() {
            this.updateList();
        },
        methods: {
            updateList: function() {
                console.log('update list');
                axios
                    .get('https://firl.us/api/shortcut')
                    .then(response => (this.list = response.data))
            },
            startLoading: function() {
                this.loading = true;
            },
            stopLoading: function() {
                this.loading = false;
            },
            displayMessage: function(msg) {
                console.log(msg);
            },
            enableCreator: function() {
                this.creator = true;
            },
            disableCreator: function() {
                this.creator = false;
            }
        }
    }
</script>
