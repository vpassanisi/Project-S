<template>
  <nav
    class="fixed z-10 top-0 left-0 w-full bg-white dark:bg-dark-gray-900 transition-nav duration-500 ease-in-out border-b border-gray-500"
    :class="[scroll === 0 ? 'h-20 py-2' : 'h-12 py-1']"
    @openLoginModal="openLoginModal"
  >
    <div class="flex flex-row items-center justify-center w-full h-full">
      <div class="flex flex-row justify-between items-center w-11/12 h-full">
        <div class="flex flex-row h-full">
          <svg
            class="dark:text-white fill-current h-full px-3"
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            x="0px"
            y="0px"
            viewBox="0 0 396.821 396.821"
            style="enable-background:new 0 0 396.821 396.821;"
            xml:space="preserve"
          >
            <g>
              <g>
                <g>
                  <path
                    d="M394.583,8.054c-0.422-3.413-3.11-6.101-6.522-6.523c-30.748-3.803-62.477-0.488-91.767,9.583c-29.293,10.072-56.355,26.973-78.258,48.876l-49.983,49.983l-72.149,9.305c-1.645,0.212-3.172,0.963-4.345,2.135l-69.563,69.563c-1.969,1.969-2.685,4.868-1.858,7.528c0.825,2.66,3.058,4.643,5.796,5.15l52.597,9.742l10.335,10.335l-22.003,11.915c-2.082,1.127-3.51,3.172-3.851,5.515s0.444,4.709,2.118,6.383l83.438,83.438c1.417,1.417,3.329,2.197,5.304,2.197c0.358,0,0.72-0.026,1.08-0.078c2.343-0.341,4.388-1.769,5.515-3.851l11.916-22.003l10.335,10.335l9.742,52.597c0.508,2.739,2.49,4.971,5.15,5.797c0.731,0.227,1.48,0.337,2.224,0.337c1.96,0,3.876-0.769,5.305-2.197l69.563-69.563c1.172-1.172,1.923-2.7,2.135-4.344l9.306-72.15l49.983-49.984c21.903-21.903,38.804-48.964,48.876-78.257C395.072,70.528,398.385,38.795,394.583,8.054z M79.674,198.355l-36.989-6.851l57.673-57.675l50.332-6.491L79.674,198.355zM152.065,313.268L82.846,244.05l17.085-9.252l61.385,61.386L152.065,313.268z M262.285,295.756l-57.674,57.674l-6.852-36.988l71.017-71.017L262.285,295.756z M325.517,167.471l-135.85,135.85l-96.874-96.874l135.85-135.851c19.738-19.739,44.002-35.076,70.287-44.49c3.395,17.492,11.948,33.719,24.654,46.424c12.705,12.706,28.931,21.259,46.424,24.655C360.593,123.468,345.255,147.732,325.517,167.471z M374.523,82.774c-15.203-2.593-29.345-9.863-40.333-20.85c-10.988-10.987-18.257-25.13-20.85-40.333c21.741-5.859,44.579-7.857,66.99-5.807C382.381,38.195,380.382,61.033,374.523,82.774z"
                  />
                  <path
                    d="M221.325,110.443c-17.74,17.741-17.74,46.606,0,64.347c8.871,8.871,20.521,13.305,32.174,13.305c11.649,0,23.304-4.436,32.173-13.305h0.001c17.74-17.74,17.74-46.606-0.001-64.347C267.931,92.703,239.065,92.704,221.325,110.443z M275.066,164.183c-11.894,11.893-31.244,11.891-43.134,0c-11.893-11.892-11.893-31.242,0-43.134c5.945-5.946,13.756-8.918,21.566-8.918c7.811,0,15.621,2.973,21.566,8.918C286.957,132.941,286.957,152.291,275.066,164.183z"
                  />
                  <path
                    d="M98.365,299.165c-2.93-2.929-7.678-2.929-10.607,0L23.41,363.512c-2.929,2.929-2.929,7.678,0,10.606c1.465,1.464,3.385,2.197,5.304,2.197s3.839-0.732,5.304-2.197l64.347-64.347C101.293,306.843,101.293,302.094,98.365,299.165z"
                  />
                  <path
                    d="M108.263,319.671l-28.991,28.991c-2.929,2.929-2.929,7.678,0,10.606c1.465,1.464,3.385,2.197,5.304,2.197s3.839-0.732,5.304-2.197l28.991-28.991c2.929-2.929,2.929-7.678,0-10.606C115.941,316.742,111.193,316.742,108.263,319.671z"
                  />
                  <path
                    d="M69.123,361.919c-3.138,0-6.002,2.024-7.062,4.973c-1.078,2.998-0.075,6.441,2.416,8.416c2.547,2.02,6.266,2.13,8.928,0.265c2.84-1.99,3.992-5.81,2.639-9.024C74.931,363.774,72.099,361.919,69.123,361.919z"
                  />
                  <path
                    d="M76.044,366.549C76.234,367,75.864,366.099,76.044,366.549L76.044,366.549z"
                  />
                  <path
                    d="M47.91,380.025l-3.992,3.992c-2.93,2.929-2.93,7.678-0.001,10.607c1.465,1.464,3.384,2.197,5.304,2.197c1.919,0,3.839-0.732,5.303-2.196l3.992-3.992c2.93-2.929,2.93-7.678,0.001-10.606C55.588,377.099,50.838,377.096,47.91,380.025z"
                  />
                  <path
                    d="M42.502,314.014c-2.93-2.929-7.678-2.929-10.607,0L2.904,343.005c-2.929,2.929-2.929,7.678,0,10.606c1.465,1.464,3.385,2.197,5.304,2.197s3.839-0.732,5.304-2.197l28.991-28.991C45.431,321.692,45.431,316.943,42.502,314.014z"
                  />
                  <path
                    d="M54.472,311.136c3.043-0.765,5.327-3.417,5.644-6.537c0.311-3.055-1.369-6.049-4.096-7.427c-2.895-1.464-6.523-0.853-8.769,1.494c-2.405,2.513-2.752,6.426-0.852,9.335c-0.06-0.09-0.106-0.156,0.015,0.029c0.126,0.185,0.083,0.118,0.023,0.029C48.204,310.626,51.429,311.901,54.472,311.136z"
                  />
                  <path
                    d="M73.867,293.257l3.991-3.992c2.929-2.929,2.929-7.678-0.001-10.606c-2.932-2.93-7.681-2.929-10.606,0.001l-3.991,3.992c-2.929,2.929-2.929,7.678,0.001,10.606c1.465,1.464,3.384,2.196,5.303,2.196C70.483,295.454,72.403,294.722,73.867,293.257z"
                  />
                </g>
              </g>
            </g>
          </svg>

          <button class="h-full mr-4 focus:outline-none">Momentum</button>

          <DropDown class="hidden md:inline-block">
            <template v-slot:button>
              <button
                class="flex flex-row items-center justify-between w-full h-full px-2 focus:outline-none shadow"
              >
                {{ routeText }}
                <i class="material-icons">arrow_drop_down</i>
              </button>
            </template>
            <template v-slot:content>
              <div
                class="absolute left-0 w-max-content bg-gray-50 dark:bg-dark-gray-800 rounded mt-1 shadow-default"
              >
                <router-link
                  v-if="isAuthenticated && $route.name !== 'Create Sub'"
                  class="flex flex-row items-center justify-start w-full p-4 no-underline hover:text-gray-400 dark:hover:text-gray-600 transition-text-color duration-200 ease-in-out"
                  :to="{ path: `/subs/create` }"
                >
                  <i class="material-icons mr-2">create</i> Create Sub
                </router-link>
                <router-link
                  v-if="
                    isAuthenticated &&
                      $route.name !== 'Create Post' &&
                      $route.params.sub
                  "
                  class="flex flex-row items-center justify-start w-full p-4 no-underline hover:text-gray-400 dark:hover:text-gray-600 transition-text-color duration-200 ease-in-out"
                  :to="{ path: `/s/${$route.params.sub}/create` }"
                >
                  <i class="material-icons mr-2">create</i> Create Post
                </router-link>
                <router-link
                  v-if="$route.name !== 'Home'"
                  class="flex flex-row items-center justify-start w-full p-4 no-underline hover:text-gray-400 dark:hover:text-gray-600 transition-text-color duration-200 ease-in-out"
                  :to="{ path: `/` }"
                >
                  <i class="material-icons mr-2">home</i>Home
                </router-link>
              </div>
            </template>
          </DropDown>
        </div>
        <div class="hidden md:flex flex-row h-full">
          <DropDown>
            <template v-slot:button>
              <button
                class="flex flex-row items-center justify-between w-full h-full px-2 focus:outline-none shadow"
              >
                <i v-if="!name" class="material-icons">person</i>
                {{ name ? name : "" }}
                <i class="material-icons">arrow_drop_down</i>
              </button>
            </template>
            <template v-slot:content>
              <div
                class="absolute right-0 w-max-content bg-gray-50 dark:bg-dark-gray-800 rounded mt-1"
              >
                <button
                  v-if="!isAuthenticated"
                  class="flex flex-row items-center justify-start w-full p-4 focus:outline-none hover:text-gray-400 dark:hover:text-gray-600 transition-text-color duration-200 ease-in-out"
                  @click="openLoginModal"
                >
                  <i class="material-icons mr-2">login</i> Login
                </button>
                <button
                  v-if="!isAuthenticated"
                  class="flex flex-row items-center justify-start w-full p-4 focus:outline-none hover:text-gray-400 dark:hover:text-gray-600 transition-text-color duration-200 ease-in-out"
                  @click="openRegisterModal"
                >
                  <i class="material-icons mr-2">person_add</i> Sign Up
                </button>
                <button
                  v-if="isAuthenticated"
                  class="flex flex-row items-center justify-start w-full p-4 focus:outline-none hover:text-gray-400 dark:hover:text-gray-600 transition-text-color duration-200 ease-in-out"
                  @click="logout"
                >
                  <i class="material-icons mr-2">west</i> Log Out
                </button>
              </div>
            </template>
          </DropDown>
          <DarkModeToggle />
        </div>
        <button
          class="flex md:hidden h-full items-center justify-center focus:outline-none px-2"
          @click="openSidebar"
        >
          <i class="material-icons">menu</i>
        </button>
      </div>
      <LoginModal v-show="loginModal" @closeModal="closeLoginModal" />
      <RegisterModal v-show="registerModal" @closeModal="closeRegisterModal" />
      <Sidebar v-show="sidebar" />
      <Error />
    </div>
    <div
      class="absolute progress-line"
      style="bottom: -3px;"
      :class="[isDataLoading || isAuthLoading ? 'visible' : 'invisible']"
    />
  </nav>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import DarkModeToggle from "../components/DarkModeToggle.vue";
import LoginModal from "../components/LoginModal.vue";
import RegisterModal from "../components/RegisterModal.vue";
import Sidebar from "../layout/Sidebar.vue";
import Error from "../layout/Error.vue";
import DropDown from "../components/DropDown.vue";

export default defineComponent({
  name: "Navbar",
  components: {
    DarkModeToggle,
    LoginModal,
    RegisterModal,
    Sidebar,
    Error,
    DropDown,
  },
  data: function() {
    return {
      routeText: "" as string | string[] | symbol,
      scroll: document.documentElement.scrollTop,
      perfersDark: window.matchMedia("(prefers-color-scheme: dark)").matches,
    };
  },
  computed: {
    isDataLoading(): boolean {
      return this.$store.direct.state.DataMod.isLoading;
    },
    name(): string {
      return this.$store.direct.state.AuthMod.name;
    },
    isAuthLoading(): boolean {
      return this.$store.direct.state.AuthMod.isAuthLoading;
    },
    isAuthenticated(): boolean {
      return this.$store.direct.state.AuthMod.isAuthenticated;
    },
    registerModal(): boolean {
      return this.$store.direct.state.EventMod.registerModal;
    },
    sidebar(): boolean {
      return this.$store.direct.state.EventMod.sidebar;
    },
    loginModal(): boolean {
      return this.$store.direct.state.EventMod.loginModal;
    },
  },
  methods: {
    logout() {
      this.$store.direct.dispatch.logout();
    },
    closeSidebar() {
      this.$store.direct.commit.closeSidebar();
    },
    openSidebar() {
      this.$store.direct.commit.openSidebar();
    },
    closeRegisterModal() {
      this.$store.direct.commit.closeRegisterModal();
    },
    openRegisterModal() {
      this.$store.direct.commit.openRegisterModal();
    },
    closeLoginModal() {
      this.$store.direct.commit.closeLoginModal();
    },
    openLoginModal() {
      this.$store.direct.commit.openLoginModal();
    },
    handleLogout() {
      this.$store.direct.commit.logout();
    },
  },
  watch: {
    $route(to, from) {
      if (to.path != from.path && to.params.sub) {
        this.routeText = to.params.sub;
      } else {
        this.routeText = to.name;
      }
    },
  },
  mounted: function() {
    document.addEventListener(
      "scroll",
      () => (this.scroll = document.documentElement.scrollTop)
    );

    if (this.$route.params.sub) {
      this.routeText = this.$route.params.sub;
    } else if (this.$route.name) {
      this.routeText = this.$route.name;
    }

    if (this.perfersDark) this.$store.direct.dispatch.turnOn();

    this.$store.direct.dispatch.me();
  },
});
</script>
