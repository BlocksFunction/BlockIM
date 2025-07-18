<script setup lang="ts">
import WelcomeComponent from "@/components/Auth/WelcomeComponent.vue";
import SwitchTheme from "@/components/SwitchTheme.vue";
import useUser from "@/stores/useUser";
import { ref } from "vue";
import { useRouter } from "vue-router";
import { toast } from "vue-sonner";
import axios from "axios";

interface RegisterAPI {
	status: string;
	clientId?: string;
	token?: string;
}

const router = useRouter();
const user = useUser();

const formValues = ref({
	username: "",
	email: "",
	password: "",
	isSubmitting: false,
});

async function handleSubmit() {
	formValues.value.isSubmitting = true;

	try {
    const response = await axios.postForm<RegisterAPI>("/api/register", {
      username: formValues.value.username,
      email: formValues.value.email,
      password: formValues.value.password,
    });

		if (response.data.token) {
			user.setUser(response.data.token);
			await router.push("/");
		}
	} catch (error: any) {
		toast.error("发生了一些错误!", {
			description: error?.data?.reason || "未知错误, 请联系网站管理员",
		});
	} finally {
		formValues.value.isSubmitting = false;
	}
}
</script>

<template>
	<div class="flex w-full h-screen overflow-hidden">
		<SwitchTheme class="fixed mt-5 right-0 mr-5" />
		<WelcomeComponent />
		<div
			class="w-full md:w-1/2 bg-gray-50 dark:bg-[#1c2735ff] flex items-center justify-center"
		>
			<div class="w-full max-w-sm mx-4 md:mx-0">
				<h1
					class="cursor-pointer text-4xl font-medium mb-8 text-center bg-gradient-to-r from-blue-700 to-blue-400 bg-clip-text text-transparent"
				>
					注册
				</h1>
				<form @submit.prevent="handleSubmit" class="space-y-6">
					<div>
						<label
							for="userid"
							class="block text-sm font-medium text-gray-700 mb-2 dark:text-gray-300"
						>
							帐号名称
						</label>
						<input
							id="userid"
							type="text"
							v-model="formValues.username"
							class="appearance-none bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg w-full py-3 px-4 text-gray-700 dark:text-white leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
							placeholder="请输入帐号名称"
							required
						/>
					</div>

					<div>
						<label
							for="email"
							class="block text-sm font-medium text-gray-700 mb-2 dark:text-gray-300"
						>
							邮箱
						</label>
						<input
							id="email"
							type="text"
							v-model="formValues.email"
							class="appearance-none bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg w-full py-3 px-4 text-gray-700 dark:text-white leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
							placeholder="请输入邮箱地址"
							required
						/>
					</div>

					<div>
						<label
							for="password"
							class="block text-sm font-medium text-gray-700 mb-2 dark:text-gray-300"
						>
							密码
						</label>
						<input
							id="password"
							type="password"
							v-model="formValues.password"
							class="appearance-none bg-white dark:bg-gray-700 border border-gray-300 dark:border-gray-600 rounded-lg w-full py-3 px-4 text-gray-700 dark:text-white leading-tight focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent"
							placeholder="请输入密码"
							required
						/>
					</div>

					<button
						type="submit"
						:disabled="formValues.isSubmitting"
						class="cursor-pointer bg-gradient-to-r from-blue-700 to-blue-500 text-white font-medium rounded-lg text-sm px-5 py-2.5 text-center w-full focus:ring-4 focus:ring-blue-300 hover:bg-blue-800 transition-colors dark:focus:ring-blue-800 disabled:opacity-50"
					>
						{{
							formValues.isSubmitting
							? "注册中..."
							: "注册"
						}}
					</button>
				</form>

				<div
					class="mt-8 text-center text-sm text-gray-500 dark:text-gray-400"
				>
					<p>
						已经有账号？
						<RouterLink
							to="/login"
							class="cursor-pointer text-blue-500 hover:text-blue-700 dark:hover:text-blue-400 ml-1"
						>
							前去登录
						</RouterLink>
					</p>
				</div>
			</div>
		</div>
	</div>
</template>
