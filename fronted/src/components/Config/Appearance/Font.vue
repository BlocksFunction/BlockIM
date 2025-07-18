<script setup lang="ts">
import { ref } from "vue";
import { Minus, Plus, TextCursorInput, Type } from "lucide-vue-next";

const fontSizes = ref([
	{ id: "xs", name: "较小", size: "text-xs" },
	{ id: "sm", name: "小", size: "text-sm" },
	{ id: "base", name: "标准", size: "text-base" },
	{ id: "lg", name: "大", size: "text-lg" },
	{ id: "xl", name: "较大", size: "text-xl" },
]);

const fontFamilies = ref([
	{ id: "sans", name: "无衬线体", class: "font-sans", preview: "Aa" },
	{ id: "serif", name: "衬线体", class: "font-serif", preview: "Aa" },
	{ id: "mono", name: "等宽体", class: "font-mono", preview: "Aa" },
]);

const selectedFontSize = ref("base");
const selectedFontFamily = ref("sans");
const lineHeight = ref(1.5);
const letterSpacing = ref(0);

const increaseFontSize = () => {
	const currentIndex = fontSizes.value.findIndex((f) =>
		f.id === selectedFontSize.value
	);
	if (currentIndex < fontSizes.value.length - 1) {
		selectedFontSize.value = fontSizes.value[currentIndex + 1].id;
	}
};

const decreaseFontSize = () => {
	const currentIndex = fontSizes.value.findIndex((f) =>
		f.id === selectedFontSize.value
	);
	if (currentIndex > 0) {
		selectedFontSize.value = fontSizes.value[currentIndex - 1].id;
	}
};
</script>

<template>
	<div class="flex-1 overflow-y-auto p-6">
		<div class="max-w-3xl mx-auto">
			<div class="mb-8">
				<h1 class="text-2xl font-bold text-gray-900 dark:text-white">
					字体设置
				</h1>
				<p class="text-gray-500 dark:text-gray-400 mt-2">
					自定义应用的字体样式和大小
				</p>
			</div>

			<div
				class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6 mb-6"
			>
				<div class="flex items-start mb-6">
					<Type class="w-6 h-6 text-blue-500 mt-1 mr-3" />
					<div class="flex-1">
						<h2
							class="text-lg font-semibold text-gray-900 dark:text-white"
						>
							字体大小
						</h2>
						<p
							class="text-sm text-gray-500 dark:text-gray-400 mt-1"
						>
							调整应用中的文字大小
						</p>
					</div>
				</div>

				<div class="flex items-center justify-between mb-4">
					<div
						class="text-sm font-medium text-gray-700 dark:text-gray-300"
					>
						当前大小: {{
							fontSizes.find((f) =>
								f.id ===
									selectedFontSize
							)?.name
						}}
					</div>
					<div class="flex space-x-2">
						<button
							@click="decreaseFontSize"
							class="p-2 rounded-lg bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600"
							:disabled='
								selectedFontSize ===
								"xs"
							'
							:class='
								selectedFontSize === "xs"
								? "opacity-50 cursor-not-allowed"
								: ""
							'
						>
							<Minus class="w-4 h-4" />
						</button>
						<button
							@click="increaseFontSize"
							class="p-2 rounded-lg bg-gray-100 dark:bg-gray-700 hover:bg-gray-200 dark:hover:bg-gray-600"
							:disabled='
								selectedFontSize ===
								"xl"
							'
							:class='
								selectedFontSize === "xl"
								? "opacity-50 cursor-not-allowed"
								: ""
							'
						>
							<Plus class="w-4 h-4" />
						</button>
					</div>
				</div>

				<div class="space-y-4">
					<div
						v-for="size in fontSizes"
						:key="size.id"
						class="p-4 border rounded-lg transition-colors"
						:class='
							[
								selectedFontSize === size.id
									? "border-blue-500 bg-blue-50 dark:bg-blue-900/20"
									: "border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700",
								size.size,
							]
						'
					>
						<div class="flex justify-between items-center">
							<div>
								<div
									class="font-medium text-gray-900 dark:text-white"
								>
									{{
										size
										.name
									}} 字体
								</div>
								<div
									class="text-gray-500 dark:text-gray-400 mt-1"
								>
									这是 {{
										size
										.name
									}} 字体的预览效果，适用于阅读体验。
								</div>
							</div>
							<button
								@click="
									selectedFontSize =
									size.id
								"
								class="px-3 py-1 text-sm rounded-lg"
								:class='
									selectedFontSize ===
										size.id
									? "bg-blue-500 text-white"
									: "bg-gray-100 dark:bg-gray-700 text-gray-700 dark:text-gray-300"
								'
							>
								选择
							</button>
						</div>
					</div>
				</div>
			</div>

			<div
				class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6 mb-6"
			>
				<div class="flex items-start mb-6">
					<Type class="w-6 h-6 text-blue-500 mt-1 mr-3" />
					<div class="flex-1">
						<h2
							class="text-lg font-semibold text-gray-900 dark:text-white"
						>
							字体系列
						</h2>
						<p
							class="text-sm text-gray-500 dark:text-gray-400 mt-1"
						>
							选择您喜欢的字体类型
						</p>
					</div>
				</div>

				<div class="grid grid-cols-1 md:grid-cols-3 gap-4">
					<button
						v-for="font in fontFamilies"
						:key="font.id"
						@click="selectedFontFamily = font.id"
						class="p-4 border rounded-xl transition-colors group"
						:class='
							[
								selectedFontFamily === font.id
									? "border-blue-500 bg-blue-50 dark:bg-blue-900/20"
									: "border-gray-200 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-700",
							]
						'
					>
						<div class="flex flex-col items-center">
							<div
								class="text-3xl mb-3"
								:class="font.class"
							>
								{{
									font
									.preview
								}}
							</div>
							<span
								class="font-medium"
								:class='
									selectedFontFamily ===
										font.id
									? "text-blue-500"
									: "text-gray-700 dark:text-gray-300"
								'
							>
								{{
									font
									.name
								}}
							</span>
							<div
								class="mt-3 text-sm text-center"
								:class="font.class"
							>
								这是一段 {{
									font
									.name
								}} 的预览文本
							</div>
						</div>
					</button>
				</div>
			</div>

			<div class="bg-white dark:bg-gray-800 rounded-xl shadow-sm p-6">
				<div class="flex items-start mb-6">
					<TextCursorInput class="w-6 h-6 text-blue-500 mt-1 mr-3" />
					<div class="flex-1">
						<h2
							class="text-lg font-semibold text-gray-900 dark:text-white"
						>
							高级设置
						</h2>
						<p
							class="text-sm text-gray-500 dark:text-gray-400 mt-1"
						>
							微调字体显示效果
						</p>
					</div>
				</div>

				<div class="space-y-6">
					<div>
						<div class="flex justify-between mb-2">
							<label
								class="block text-sm font-medium text-gray-700 dark:text-gray-300"
							>行高</label>
							<span
								class="text-sm text-gray-500 dark:text-gray-400"
							>{{
								lineHeight
								.toFixed(1)
							}}</span>
						</div>
						<input
							v-model.number="lineHeight"
							type="range"
							min="1.0"
							max="2.0"
							step="0.1"
							class="w-full h-2 bg-gray-200 dark:bg-gray-700 rounded-lg appearance-none cursor-pointer accent-blue-500"
						/>
						<div
							class="mt-4 p-4 bg-gray-50 dark:bg-gray-900 rounded-lg"
						>
							<p
								:style="
									{
										lineHeight:
											lineHeight,
									}
								"
								class="transition-all"
							>
								行高设置影响文本行与行之间的垂直间距。适当的行高可以提升阅读体验，使文本更易读。当前行高为
								{{
									lineHeight
									.toFixed(
										1,
									)
								}} 倍字体大小。
							</p>
						</div>
					</div>

					<div>
						<div class="flex justify-between mb-2">
							<label
								class="block text-sm font-medium text-gray-700 dark:text-gray-300"
							>字间距</label>
							<span
								class="text-sm text-gray-500 dark:text-gray-400"
							>{{ letterSpacing }}px</span>
						</div>
						<input
							v-model.number="letterSpacing"
							type="range"
							min="0"
							max="5"
							step="0.5"
							class="w-full h-2 bg-gray-200 dark:bg-gray-700 rounded-lg appearance-none cursor-pointer accent-blue-500"
						/>
						<div
							class="mt-4 p-4 bg-gray-50 dark:bg-gray-900 rounded-lg"
						>
							<p
								:style="
									{
										letterSpacing:
											`${letterSpacing}px`,
									}
								"
								class="transition-all"
							>
								字间距设置影响字符之间的水平距离。适当的字间距可以提升文本的可读性和美观度。当前字间距为
								{{ letterSpacing }} 像素。
							</p>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>
