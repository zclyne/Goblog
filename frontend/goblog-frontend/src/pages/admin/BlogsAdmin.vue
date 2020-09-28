<template>
    <q-page padding>
        <div class="blogs-admin-container">
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
                            <q-btn flat dense round size="12px" icon="delete" @click="deleteBlog(blogView)"></q-btn>
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
    export default {
        name: 'BlogsAdmin',
        data () {
            return {
                blogViewList: []
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
            deleteBlog (blogView) {
                const blogId = blogView.blog.blog_id
                axios.delete('/api/blogs/' + blogId)
                    .then(this.refreshBlogList)
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
