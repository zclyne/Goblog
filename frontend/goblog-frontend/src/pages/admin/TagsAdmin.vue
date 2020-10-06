<template>
    <q-page padding>
        <div class="tags-admin-container">
            <message-banner :message="message" :showMessage="showMessage" />
            <q-list bordered class="rounded-borders shadow-3">
                <q-item-label header>Create New Tag</q-item-label>
                <q-item>
                    <q-input
                            style="width: 100%"
                            v-model="newTagName"
                            label="Tag Name"
                            hint="1~20 characters">
                        <template v-slot:append>
                            <q-btn flat dense round size="12px" icon="add" @click="createNewTag"></q-btn>
                        </template>
                    </q-input>

                </q-item>
                <q-separator spaced/>
                <q-item-label header>Tags</q-item-label>
                <q-item v-for="tag in tagList" :key="tag.tag_id">
                    <q-item-section>
                        {{tag.name}}
                    </q-item-section>
                    <q-item-section side>
                        <div>
                            <q-btn flat dense round size="12px" icon="edit">
                                <q-popup-edit
                                        v-model="tag.name"
                                        :validate="tagName => {return tagName.length > 0 && tagName.length <= 20}"
                                        @save="editTagName(tag)">
                                    <q-input
                                            v-model="tag.name"
                                            dense
                                            autofocus
                                            counter
                                            hint="1~20 chars"></q-input>
                                </q-popup-edit>
                            </q-btn>
                            <q-btn flat dense round size="12px" icon="delete" @click="deleteTag(tag)"></q-btn>
                        </div>
                    </q-item-section>
                </q-item>
            </q-list>
        </div>
    </q-page>
</template>

<script>
    import axios from 'axios'
    import MessageBanner from 'components/MessageBanner.vue'
    export default {
        name: 'TagsAdmin',
        components: {
            MessageBanner
        },
        data () {
            return {
                newTagName: '',
                tagList: [],
                message: '',
                showMessage: false
            }
        },
        methods: {
            refreshTagList () {
                axios.get('/api/tags')
                    .then(this.listTagsSuccess)
            },
            listTagsSuccess (res) {
                if (res.data) {
                    this.tagList = res.data
                }
            },
            createNewTag () {
                axios.post('/api/tags', {
                    name: this.newTagName
                }).then(() => {
                    this.newTagName = ''
                    this.setMessageAndRefresh('Successfully created the new tag')
                })
            },
            editTagName (tag) {
                axios.put('/api/tags', tag)
                    .then(this.refreshTagList)
            },
            deleteTag (tag) {
                axios.delete('/api/tags/' + tag.tag_id)
                    .then(() => {
                        this.setMessageAndRefresh('Successfully deleted the tag')
                    })
            },
            setMessageAndRefresh (msg) {
                this.message = msg
                this.showMessage = true
                this.refreshTagList()
                window.clearTimeout(this.timer)
                this.timer = window.setTimeout(() => {this.showMessage = false}, 3000)
            }
        },
        mounted () {
            this.refreshTagList()
        }
    }
</script>

<style scoped>
    .tags-admin-container {
        margin: 0 auto;
        max-width: 1000px;
    }
</style>
