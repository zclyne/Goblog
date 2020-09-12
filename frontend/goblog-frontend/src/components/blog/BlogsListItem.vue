<template>
    <q-card class="blog-list-item-card q-mt-md q-mx-auto shadow-5" @click="handleCardClick">
        <q-card-section horizontal>
            <q-img ratio="4/3" class="col-3 blog-img" :src="blog.image_url" alt="" />
            <q-card-section class="col-9">
                <!-- blog title and type area -->
                <q-card-section class="text-h5" horizontal>
                    <q-card-section class="col-9">
                        {{blog.title}}
                    </q-card-section>
                    <q-card-section class="col-3">
                        <q-chip outline color="secondary float-right" text-color="white" icon="bookmark">
                            {{type.name}}
                        </q-chip>
                    </q-card-section>
                </q-card-section>
                <!-- blog content area -->
                <q-card-section class="text-body1">
                    {{blog.content}}
                </q-card-section>
                <!-- tags of this blog -->
                <q-card-section horizontal>
                    <q-card-section class="col-8">
                        <q-chip outline color="deep-purple-4" text-color="white" icon="label"
                                v-for="tag in tags" :key="tag.tag_id">
                            {{tag.name}}
                        </q-chip>
                    </q-card-section>
                    <q-card-section class="col-4 text-body2 text-right blog-date">
                        {{date}}
                    </q-card-section>
                </q-card-section>
            </q-card-section>
        </q-card-section>
    </q-card>
</template>

<script>
    export default {
        name: 'BlogsListItem',
        props: {
            blog: Object,
            type: Object,
            tags: Array
        },
        data () {
            return {
                date: String
            }
        },
        methods: {
            handleCardClick () {
                this.$router.push('/blogs/' + this.blog.blog_id)
            }
        },
        mounted () {
            // convert date format
            this.date = new Date(this.blog.create_at).toLocaleDateString()
        }
    }
</script>

<style scoped>
    .blog-list-item-card {
        width: 100%;
        max-width: 1000px;
        max-height: 300px;
        overflow: hidden;
        text-decoration: none;
    }
    .blog-img {
        max-height: 100%;
    }
    .blog-date {
        margin: auto;
    }
</style>
