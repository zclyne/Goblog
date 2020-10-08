<template>
    <q-page padding>
        <div class="tags-admin-container">
            <message-banner :message="message" :color="messageBannerColor" v-if="showMessage">
                <template v-slot:action>
                    <q-btn-group v-if="deleteBtnGroupShow" flat>
                        <q-btn flat color="white" label="CONFIRM" @click="confirmDelete" />
                        <q-btn flat color="white" label="CANCEL" @click="cancelDelete" />
                    </q-btn-group>
                    <q-btn flat color="white" label="GOT IT" v-else @click="closeMessageBox" />
                </template>
            </message-banner>
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
                            <q-btn flat dense round size="12px" icon="delete" @click="deleteTagRequest(tag)"></q-btn>
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
                showMessage: false,
                messageBannerColor: '',
                deleteBtnGroupShow: false,
                tagToDelete: null,
                timer: null
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
                if (!this.validateTagName(this.newTagName)) {
                    return
                }
                axios.post('/api/tags', {
                    name: this.newTagName
                }).then(() => {
                    this.setMessageAndRefresh('Successfully created the tag')
                    this.newTagName = ''
                })
            },
            editTagName (tag) {
                if (!this.validateTagName(tag.name)) {
                    return
                }
                axios.put('/api/tags', tag)
                    .then(this.setMessageAndRefresh('Successfully edited the tag'))
            },
            deleteTagRequest (tag) {
                window.clearTimeout(this.timer)
                this.deleteBtnGroupShow = true
                this.tagToDelete = tag
                this.setMessageAndColor('Are you sure to delete this tag?', 'bg-warning')
                
            },
            confirmDelete () {
                this.showMessage = false
                this.deleteBtnGroupShow = false
                axios.delete('/api/tags/' + this.tagToDelete.tag_id)
                    .then(() => {
                        this.tagToDelete = null
                        this.setMessageAndRefresh('Successfully deleted the tag')
                    })
            },
            cancelDelete () {
                this.showMessage = false
                this.deleteBtnGroupShow = false
            },
            setMessageAndRefresh (msg) {
                this.setMessageAndColor(msg, 'bg-positive')
                window.clearTimeout(this.timer)
                this.timer = window.setTimeout(() => {this.showMessage = false}, 3000)
                this.refreshTagList()
            },
            validateTagName (tagName) {
                if (tagName.length > 0 && tagName.length <= 20) {
                    return true
                } else {
                    this.setMessageAndColor('Tag name should be of 1 ~ 20 characters', 'bg-red')
                    return false
                }
            },
            setMessageAndColor (msg, color) {
                this.message = msg
                this.messageBannerColor = color
                this.showMessage = true
            },
            closeMessageBox () {
                this.showMessage = false
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
