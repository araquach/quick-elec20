<template>
    <b-navbar class="navbar is-primary is-spaced">
        <template slot="brand">
            <h1>Quick-Elec</h1>
            <h1>North West Electricians</h1>
            <b-navbar-item href="/">
                <img src="/dist/img/qe_logo.svg" width="130">
            </b-navbar-item>
        </template>
        <template slot="start">
            <b-navbar-item href="/">
                Home
            </b-navbar-item>
            <b-navbar-item href="/about">
                About
            </b-navbar-item>
            <b-navbar-dropdown label="More">
                <b-navbar-item href="/domestic">
                    Domestic
                </b-navbar-item>
                <b-navbar-item href="/commercial">
                    Commercial
                </b-navbar-item>
                <b-navbar-item href="/exterior">
                    Exterior
                </b-navbar-item>
                <b-navbar-item href="/testimonials">
                    Testimonials
                </b-navbar-item>
            </b-navbar-dropdown>
        </template>
        <template slot="end">
            <b-navbar-item tag="div">
                <div class="buttons">
                    <a href="/contact" class="button is-light">
                        Contact Us
                    </a>
                </div>
            </b-navbar-item>
        </template>
    </b-navbar>
</template>