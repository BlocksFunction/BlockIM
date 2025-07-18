<script setup lang="ts">
import Theme from "@/components/Config/Appearance/Theme.vue";
import {
	Bell,
	Brush,
	ChevronRight,
	HelpCircle,
	IdCard,
	Inbox,
	Info,
	LogOut,
	Mail,
	Palette,
	Phone,
	PhoneCall,
	Send,
	Settings,
	ShieldUser,
	Type,
	User,
	UserLock,
} from "lucide-vue-next";
import { ref } from "vue";
import Profile from "@/components/Config/Accout/Profile.vue";
import Security from "@/components/Config/Accout/Security.vue";
import Privacy from "@/components/Config/Accout/Privacy.vue";
import Font from "@/components/Config/Appearance/Font.vue";
import FAQ from "@/components/Config/Help/FAQ.vue";
import Feedback from "@/components/Config/Help/Feedback.vue";
import Contact from "@/components/Config/Help/Contact.vue";
import About from "@/components/Config/Help/About.vue";
import Message from "@/components/Config/Notifications/Message.vue";
import Call from "@/components/Config/Notifications/Call.vue";
import System from "@/components/Config/Notifications/System.vue";

const settings = [
	{
		id: "account",
		icon: User,
		label: "账号设置",
		options: [
			{ id: "profile", label: "个人资料", icon: IdCard },
			{ id: "security", label: "安全设置", icon: ShieldUser },
			{ id: "privacy", label: "隐私设置", icon: UserLock },
		],
	},
	{
		id: "notifications",
		icon: Bell,
		label: "通知设置",
		options: [
			{ id: "message", label: "消息通知", icon: Mail },
			{ id: "call", label: "通话通知", icon: PhoneCall },
			{ id: "system", label: "系统通知", icon: Inbox },
		],
	},
	{
		id: "appearance",
		icon: Brush,
		label: "外观设置",
		options: [
			{ id: "theme", label: "主题设置", icon: Palette },
			{ id: "font", label: "字体设置", icon: Type },
		],
	},
	{
		id: "Help",
		icon: HelpCircle,
		label: "帮助与反馈",
		options: [
			{ id: "faq", label: "常见问题", icon: HelpCircle },
			{ id: "feedback", label: "意见反馈", icon: Send },
			{ id: "contact", label: "联系我们", icon: Phone },
			{ id: "about", label: "关于", icon: Info },
		],
	},
];

const expandedSection = ref<string | null>(null);

const currentPage = ref("default");

const toggleSection = (sectionId: string) => {
	expandedSection.value = expandedSection.value === sectionId
		? null
		: sectionId;
};
</script>

<template>
	<div
		class="w-full h-screen flex overflow-hidden bg-gray-50 dark:bg-gray-900"
	>
		<div
			class="w-64 h-full bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700 flex flex-col"
		>
			<div class="p-6 border-b border-gray-200 dark:border-gray-700">
				<h1 class="text-xl font-bold text-gray-800 dark:text-white">
					设置
				</h1>
				<p class="text-sm text-gray-500 dark:text-gray-400 mt-1">
					自定义您的应用体验
				</p>
			</div>

			<nav class="flex-1 overflow-y-auto p-4">
				<ul class="space-y-1">
					<li v-for="item in settings" :key="item.id">
						<button
							@click="toggleSection(item.id)"
							class="w-full flex items-center justify-between p-3 rounded-lg text-left transition-all duration-200 hover:bg-gray-100 dark:hover:bg-gray-700 group"
						>
							<span class="flex items-center">
								<component
									:is="
										item
										.icon
									"
									class="w-5 h-5 text-gray-500 dark:text-gray-400 mr-3 transition-colors group-hover:text-blue-500"
								/>
								<span
									class="font-medium text-gray-700 dark:text-gray-300"
								>
									{{
										item
										.label
									}}
								</span>
							</span>
							<ChevronRight
								class="w-4 h-4 text-gray-400 transition-transform duration-200"
								:class='
									{
										"rotate-90":
											expandedSection ===
												item
													.id,
									}
								'
							/>
						</button>

						<ul
							v-if="
								expandedSection ===
								item.id
							"
							class="ml-10 mt-1 space-y-1 animate-accordion-down"
						>
							<li v-for="option in item.options" :key="option.id">
								<button
									class="w-full flex items-center justify-between p-2 rounded-lg text-sm text-left text-gray-600 dark:text-gray-400 hover:bg-gray-100 dark:hover:bg-gray-700 transition-colors"
									:class='
										currentPage ===
											option
												.id
										? "bg-gray-200 dark:bg-gray-600"
										: ""
									'
									@click="
										currentPage =
										option
											.id
									"
								>
									<span class="flex items-center">
										<component
											:is="
												option
												.icon
											"
											class="w-5 h-5 text-gray-500 dark:text-gray-400 mr-3 transition-colors"
										/>
										<span
											class="font-medium text-gray-700 dark:text-gray-300"
										>
											{{
												option
												.label
											}}
										</span>
									</span>

									<ChevronRight
										class="w-4 h-4 opacity-0 hover:opacity-100 transition-opacity"
									/>
								</button>
							</li>
						</ul>
					</li>
				</ul>
			</nav>

			<div
				class="p-4 border-t border-gray-200 dark:border-gray-700 mt-auto"
			>
				<button
					class="w-full flex items-center justify-center space-x-2 p-3 text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20 rounded-lg transition-colors"
				>
					<LogOut class="w-5 h-5" />
					<span>退出登录</span>
				</button>
			</div>
		</div>

		<template
			v-if='currentPage == "default"'
			class="w-full h-full bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center mb-4"
		>
			<div
				class="flex-1 flex flex-col items-center justify-center p-8 text-center"
			>
				<div
					class="w-32 h-32 bg-gray-100 dark:bg-gray-800 rounded-full flex items-center justify-center mb-6"
				>
					<Settings
						class="text-gray-400 dark:text-gray-500"
						:size="48"
					/>
				</div>
				<h2
					class="text-2xl font-semibold text-gray-700 dark:text-gray-300 mb-2"
				>
					设置
				</h2>
				<p class="text-gray-500 dark:text-gray-400 max-w-md">
					从左侧自定义信息，快来展现你的创造力吧！
				</p>
			</div>
		</template>
		<template v-if='currentPage == "profile"'>
			<Profile />
		</template>
		<template v-if='currentPage == "security"'>
			<Security />
		</template>
		<template v-if='currentPage == "privacy"'>
			<Privacy />
		</template>
		<template v-if='currentPage == "theme"'>
			<Theme />
		</template>
		<template v-if='currentPage == "font"'>
			<Font />
		</template>
		<template v-if='currentPage == "faq"'>
			<FAQ />
		</template>
		<template v-if='currentPage == "feedback"'>
			<Feedback />
		</template>
		<template v-if='currentPage == "contact"'>
			<Contact />
		</template>
		<template v-if='currentPage == "about"'>
			<About />
		</template>
		<template v-if='currentPage == "message"'>
			<Message />
		</template>
		<template v-if='currentPage == "call"'>
			<Call />
		</template>
		<template v-if='currentPage == "system"'>
			<System />
		</template>
	</div>
</template>
