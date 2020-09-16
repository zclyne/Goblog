<template>
    <q-page padding>
        <div class="tags-admin-container">
            <q-list bordered class="rounded-borders shadow-3">
                <q-item-label header>Create New Tag</q-item-label>
                <q-item>
                    <q-input
                            style="width: 100%"
                            v-model="newTagName"
                            label="Enter New Tag Name with 1~20 chars">
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
    export default {
        name: 'TagsAdmin',
        data () {
            return {
                newTagName: '',
                tagList: []
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
                    this.refreshTagList()
                })
            },
            editTagName (tag) {
                axios.put('/api/tags', tag)
                    .then(this.refreshTagList)
            },
            deleteTag (tag) {
                axios.delete('/api/tags/' + tag.tag_id)
                    .then(this.refreshTagList)
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
