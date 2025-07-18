<script setup lang="ts">
import useTheme from "@/stores/useTheme.ts";
import { MoonStar, Sun } from "lucide-vue-next";

const theme = useTheme();
function handleThemeToggle() {
	theme.toggleTheme();
}
</script>

<template>
	<button @click="handleThemeToggle" class="cursor-pointer">
		<Sun
			:size="30"
			class="text-yellow-500"
			v-if='theme.getTheme == "light"'
		/>
		<MoonStar class="text-blue-500" :size="30" v-else />
	</button>
</template>
