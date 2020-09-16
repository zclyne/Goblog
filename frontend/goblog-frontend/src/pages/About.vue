<template>
    <q-page padding>
        <q-card class="about-card">
            <q-card-section class="text-h4 q-pa-md" horizontal>
                <div class="col-2">
                    <q-avatar size="100px" class="q-mb-sm">
                        <img src="../static/avatar.jpg">
                    </q-avatar>
                </div>
                <div class="col-auto">
                    <div class="about-title vertical-middle">Yifan Zhang</div>
                </div>
            </q-card-section>
            <q-separator />
            <q-card-section class="text-h6">
                Bachelor of Engineering in Electronic Engineering@Shanghai Jiao Tong University
                <br>
                Master of Science in Computer Science@Columbia University
                <br>
                Software Developer Intern@VMware
                <br>
                Software Developer Intern@Microsoft
            </q-card-section>
        </q-card>
    </q-page>
</template>

<script>
    export default {
        name: 'About'
    }
</script>

<style scoped>
    .about-card {
        max-width: 1000px;
        margin: 0 auto;
    }
    .about-title {
        margin: auto 0;
    }
</style>
