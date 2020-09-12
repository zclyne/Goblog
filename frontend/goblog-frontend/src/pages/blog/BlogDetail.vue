<template>
    <q-page padding>
        <div class="column blog-detail-container q-mt-md q-mx-auto">
            <div class="row">
                <div class="col-9">
                    <div class="text-h4">
                        {{blog.title}}
                    </div>
                    <div class="text-grey-9">
                        {{date}}
                    </div>
                </div>
                <div class="col-3">
                    <q-chip outline color="secondary float-right" text-color="white" icon="bookmark">
                        {{type.name}}
                    </q-chip>
                </div>
            </div>
            <div>
                <q-img :src="blog.image_url"></q-img>
            </div>
            <div>
                <mavon-editor
                        style="width: 100%"
                        :editable="false"
                        :subfield="false"
                        defaultOpen="preview"
                        :toolbarsFlag="false"
                        :scrollStyle="true"
                        :ishljs="true"
                        :boxShadow="false"
                        :value="blog.content" />
            </div>
            <div class="row q-pt-md">
                <q-chip outline color="deep-purple-4"
                        text-color="white"
                        icon="label"
                        v-for="tag in tags" :key="tag.tag_id">
                    {{tag.name}}
                </q-chip>
            </div>

            <!-- comment area -->
            <q-list bordered padding class="q-mt-md rounded-borders">
                <div v-for="(comment, index) in comments" :key="comment.comment_id">
                    <q-item>
                        <q-item-section top avatar>
                            <q-avatar>
                                <img :src="comment.avatar_url" alt="">
                            </q-avatar>
                        </q-item-section>

                        <q-item-section>
                            <q-item-label>{{comment.nickname}}</q-item-label>
                            <q-item-label caption>{{comment.content}}</q-item-label>
                        </q-item-section>

                        <q-item-section side top>
                            <q-item-label caption>{{new Date(comment.create_at).toLocaleString()}}</q-item-label>
                        </q-item-section>
                    </q-item>
                    <q-separator spaced inset="item" v-if="index < comments.length - 1" />
                </div>
                <q-separator spaced />
                <!-- add comment area -->
                <div class="q-mt-md q-mx-md">
                    <div class="text-h6">Leave Your Comment Here</div>
                    <q-input label="Nickname" v-model="commenterNickname"></q-input>
                    <q-input label="Avatar Url" v-model="commenterAvatarUrl"></q-input>
                    <q-input label="Email" v-model="commenterEmail"></q-input>
                    <q-input label="Your Thoughts" type="textarea" v-model="commentContent"></q-input>
                    <q-btn color="primary" label="Submit" class="float-right q-mt-md" @click="submitCommentBtnClick"/>
                </div>
            </q-list>


        </div>
    </q-page>
</template>

<script>
    import axios from 'axios'
    export default {
        name: 'BlogDetail',
        data () {
            return {
                blog: '',
                type: '',
                tags: [],
                date: '',
                comments: [],
                commenterNickname: '',
                commenterEmail: '',
                commenterAvatarUrl: '',
                commentContent: ''
            }
        },
        methods: {
            getBlogByIdSuccess (res) {
                if (res.data) {
                    this.blog = res.data.blog
                    this.type = res.data.type
                    this.tags = res.data.tags
                    // convert date format
                    this.date = new Date(this.blog.create_at).toLocaleDateString()
                }
            },
            listCommentsByBlogIdSuccess (res) {
                console.log(res)
                if (res.data) {
                    this.comments = res.data
                }
            },
            submitCommentBtnClick () {
                const submitCommentRequest = {
                    blog_id: this.blog.blog_id,
                    nickname: this.commenterNickname,
                    email: this.commenterEmail,
                    avatar_url: this.commenterAvatarUrl,
                    content: this.commentContent,
                    create_at: new Date()
                }
                axios.post('/api/comments', submitCommentRequest)
                    .then(this.refreshComments)
            },
            refreshComments () {
                axios.get('/api/comments?blog_id=' + this.$route.params.id)
                    .then(this.listCommentsByBlogIdSuccess)
                // clear input area
                this.commenterNickname = ''
                this.commenterAvatarUrl = ''
                this.commenterAvatarUrl = ''
                this.commentContent = ''
            }
        },
        mounted () {
            axios.get('/api/blogs/' + this.$route.params.id)
                .then(this.getBlogByIdSuccess)
            this.refreshComments()
        }
    }
</script>

<style scoped>
    .blog-detail-container {
        max-width: 1000px;
    }
</style>
