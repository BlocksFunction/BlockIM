<script setup lang="ts">
import useTheme from "@/stores/useTheme.ts";
import { RouterView } from "vue-router";
import { Toaster } from "vue-sonner";

const theme = useTheme();
</script>

<template>
	<div :class="theme.getTheme">
		<RouterView />
		<Toaster />
	</div>
</template>
