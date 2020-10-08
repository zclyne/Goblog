<template>
    <q-page padding>
        <div class="blogs-admin-container">
            <message-banner :message="message" :color="messageBannerColor" v-if="showMessage">
                <template v-slot:action>
                    <q-btn flat color="white" label="CONFIRM" @click="confirmDelete" />
                    <q-btn flat color="white" label="CANCEL" @click="cancelDelete" />
                </template>
            </message-banner>
            <q-list bordered class="rounded-borders shadow-3">
                <q-item-label header>
                    Blogs
                </q-item-label>
                <q-item v-for="blogView in blogViewList" :key="blogView.blog_id">
                    <q-item-section>
                        {{blogView.blog.title}}
                    </q-item-section>
                    <q-item-section side>
                        <div>
                            <q-btn flat dense round size="12px" icon="edit" @click="editBlog(blogView)"></q-btn>
                            <q-btn flat dense round size="12px" icon="delete" @click="deleteBlogRequest(blogView)"></q-btn>
                        </div>
                    </q-item-section>
                </q-item>
            </q-list>
            <q-page-sticky position="bottom-right" :offset="[18, 18]">
                <q-btn fab icon="add" color="primary" @click="createBlog" />
            </q-page-sticky>
        </div>
    </q-page>
</template>

<script>
    import axios from 'axios'
    import MessageBanner from 'components/MessageBanner.vue'
    export default {
        name: 'BlogsAdmin',
        components: {
            MessageBanner
        },
        data () {
            return {
                blogViewList: [],
                messsage: '',
                showMessage: false,
                blogIdToDelete: '',
                timer: null
            }
        },
        methods: {
            refreshBlogList () {
                axios.get('/api/blogs')
                    .then(this.listBlogsSuccess)
            },
            listBlogsSuccess (res) {
                if (res.data) {
                    this.blogViewList = res.data
                    console.log(this.blogViewList)
                }
            },
            createBlog () {
                this.$router.push('/admin/blogs/edit/')
            },
            editBlog (blogView) {
                this.$router.push('/admin/blogs/edit/' + blogView.blog.blog_id)
            },
            deleteBlogRequest (blogView) {
                window.clearTimeout(this.timer)
                this.blogIdToDelete = blogView.blog.blog_id
                this.setMessageAndColor('Are you sure to delete this blog?', 'bg-warning')
            },
            confirmDelete () {
                this.showMessage = false
                axios.delete('/api/blogs/' + this.blogIdToDelete)
                    .then(this.setMessageAndRefresh('Successfully deleted the blog'))
            },
            cancelDelete () {
                this.showMessage = false
            },
            setMessageAndColor (msg, color) {
                this.message = msg
                this.messageBannerColor = color
                this.showMessage = true
            },
            setMessageAndRefresh (msg) {
                this.setMessageAndColor(msg, 'bg-primary')
                this.refreshBlogList()
                window.clearTimeout(this.timer)
                this.timer = window.setTimeout(() => {this.showMessage = false}, 3000)
            }
        },
        mounted () {
            this.refreshBlogList()
        }
    }
</script>

<style scoped>
    .blogs-admin-container {
        margin: 0 auto;
        max-width: 1000px;
    }
</style>
