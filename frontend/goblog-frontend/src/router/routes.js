import MainLayout from 'layouts/MainLayout'
import Blogs from 'pages/blog/Blogs'
import BlogDetail from 'pages/blog/BlogDetail'
import Types from 'pages/Types'
import Tags from 'pages/Tags'

const routes = [
    {
        path: '/',
        component: MainLayout,
        children: [{
            path: '/blogs',
            component: Blogs
        }, {
            path: '/blogs/:id',
            component: BlogDetail
        }, {
            path: '/types',
            component: Types
        }, {
            path: '/tags',
            component: Tags
        }
        ]
    },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '*',
    component: () => import('pages/Error404.vue')
  }
]

export default routes
