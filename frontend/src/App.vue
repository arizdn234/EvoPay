<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<template>
  <div class="app-container">
    <header>
      <div class="wrapper">
        <nav>
          <RouterLink to="/">API Docs</RouterLink>
          <RouterLink to="/users/login">Login</RouterLink>
          <RouterLink to="/users/register">Register</RouterLink>
        </nav>
      </div>
    </header>
    <main>
      <RouterView />
    </main>
  </div>
</template>

<style scoped>
/* Resetting default styles for 'a' tags */
a {
  text-decoration: none;
  color: inherit;
}

.app-container {
  display: flex;
  flex-direction: column;
}

header {
  background: #505050;
  line-height: 1.5;
  padding: 0 1rem;
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 1px solid #646464;
}

nav {
  display: flex;
  gap: 1rem;
  font-size: 1rem;
  margin: .3em 8em;
}

nav a.router-link-exact-active {
  color: white;
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

main {
  flex: 1;
  padding: 1rem;
}

@media (min-width: 1024px) {
  header {
    padding-right: calc(var(--section-gap) / 2);
  }
}
</style>
