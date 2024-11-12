<template>
  <!-- Controls -->
  <div class="flex items-center justify-center gap-8 mt-8">
    <button class="text-gray-400 hover:text-white transition-colors">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-8 w-8"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path d="M6 6h2v12H6zm3.5 6l8.5 6V6z" />
      </svg>
    </button>
    <button
      class="w-16 h-16 rounded-full bg-white hover:bg-gray-200 transition-colors flex items-center justify-center"
      @click="isPlaying = !isPlaying"
    >
      <svg
        v-if="!isPlaying"
        xmlns="http://www.w3.org/2000/svg"
        class="h-8 w-8"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path d="M8 5v14l11-7z" />
      </svg>
      <svg
        v-else
        xmlns="http://www.w3.org/2000/svg"
        class="h-8 w-8"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path d="M6 19h4V5H6v14zm8-14v14h4V5h-4z" />
      </svg>
    </button>
    <button class="text-gray-400 hover:text-white transition-colors">
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-8 w-8"
        viewBox="0 0 24 24"
        fill="currentColor"
      >
        <path d="M6 18l8.5-6L6 6v12zM16 6v12h2V6h-2z" />
      </svg>
    </button>
  </div>
</template>
