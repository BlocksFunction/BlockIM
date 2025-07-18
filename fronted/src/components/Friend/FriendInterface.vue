<script setup lang="ts">
import { debounce } from "lodash-es";
import {
	ChevronDown,
	ChevronRight,
	Plus,
	Search,
	UserPlus,
} from "lucide-vue-next";
import { computed, ref, shallowRef } from "vue";

const activeTab = ref("friends");
const activePage = ref("");
const searchQuery = ref("");

const tabs = ref([
	{ id: "friends", label: "好友" },
	{ id: "groups", label: "群聊" },
]);

const friendGroups = shallowRef([
	{
		id: "default",
		name: "默认分组",
		collapsed: false,
		items: [
			{
				id: 1,
				name: "张三",
				avatar: "张",
				online: true,
				status: "在线",
			},
			{
				id: 2,
				name: "李四",
				avatar: "李",
				online: false,
				status: "离线",
			},
			{
				id: 3,
				name: "王五",
				avatar: "王",
				online: true,
				status: "忙碌",
			},
		],
	},
	{
		id: "work",
		name: "工作分组",
		collapsed: true,
		items: [
			{
				id: 4,
				name: "陈工",
				avatar: "陈",
				online: true,
				status: "在线",
			},
		],
	},
]);

const groupCategories = shallowRef([
	{
		id: "work",
		name: "工作群组",
		collapsed: false,
		items: [
			{
				id: 1,
				name: "测试群一号",
				avatar: "测试",
				members: 15,
				lastMessage: "虽然是假的",
			},
			{
				id: 2,
				name: "测试群二号",
				avatar: "PD",
				members: 8,
				lastMessage: "虽然这也是假的",
			},
		],
	},
	{
		id: "social",
		name: "社交群组",
		collapsed: true,
		items: [
			{
				id: 3,
				name: "校友群",
				avatar: "AL",
				members: 32,
				lastMessage: "周末聚会报名开始",
			},
		],
	},
]);

const friendRequests = shallowRef([
	{
		id: 1,
		from: {
			id: 101,
			name: "张三",
			avatar: "张",
			online: true,
		},
		message: "我们一起工作吧",
		time: new Date(Date.now() - 1),
		status: "pending",
	},
	{
		id: 2,
		from: {
			id: 102,
			name: "李四",
			avatar: "李",
			online: false,
		},
		message: "",
		time: new Date(Date.now() - 86400000),
		status: "pending",
	},
]);

const friendRequestsMap = new Map(friendRequests.value.map((r) => [r.id, r]));

const formatTime = /*#__PURE__*/ (date: Date): string => {
	const now = new Date();
	const diff = now.getTime() - date.getTime();

	if (diff < 1000) return "刚刚";
	if (diff < 60000) return `${Math.floor(diff / 1000)}秒前`;
	if (diff < 3600000) return `${Math.floor(diff / 60000)}分钟前`;
	if (diff < 86400000) return `${Math.floor(diff / 3600000)}小时前`;
	if (diff < 31536000000) return `${Math.floor(diff / 86400000)}天前`;
	return `${Math.floor(diff / 31536000000)}年前`;
};

const handleRequest = (id: number, action: string) => {
	const index = friendRequestsMap.get(id);

	if (index) {
		let updateRequests = [...friendRequests.value];
		const updateFriendGroups = [...friendGroups.value];

		if (action == "accept") {
			updateFriendGroups[0].items.push({
				id: index.from.id,
				name: index.from.name,
				avatar: index.from.avatar,
				online: index.from.online,
				status: index.from.online ? "在线" : "离线",
			});
		}
		updateRequests[index.id - 1].status = action;

		friendRequests.value = updateRequests;
		friendGroups.value = updateFriendGroups;
	}
};

const filteredFriendGroups = computed(() => {
	const query = searchQuery.value.toLowerCase();
	if (!query) return friendGroups.value;
	return friendGroups.value
		.map((group) => {
			const filteredItems = group.items.filter((friend) =>
				friend.name.toLowerCase().includes(query)
			);
			return { ...group, items: filteredItems };
		})
		.filter((group) => group.items.length > 0);
});

const filteredGroupCategories = computed(() => {
	const query = searchQuery.value.toLowerCase();
	if (!query) return groupCategories.value;

	return groupCategories.value
		.map((category) => {
			const filteredItems = category.items.filter((group) =>
				group.name.toLowerCase().includes(query)
			);
			return { ...category, items: filteredItems };
		})
		.filter((category) => category.items.length > 0);
});

const toggleGroup = (groupId: string) => {
	const groupIndex = friendGroups.value.findIndex((g) => g.id === groupId);
	if (groupIndex > -1) {
		friendGroups.value[groupIndex].collapsed = !friendGroups
			.value[groupIndex].collapsed;
		friendGroups.value = [...friendGroups.value];
	}
};

const toggleCategory = (categoryId: string) => {
	const categoryIndex = groupCategories.value.findIndex(
		(c) => c.id === categoryId,
	);

	if (categoryIndex > -1) {
		groupCategories.value[categoryIndex].collapsed = !groupCategories
			.value[categoryIndex].collapsed;
		groupCategories.value = [...groupCategories.value];
	}
};

const truncateText = (text: string, length: number, useEllipsis = true) => {
	if (text.length <= length) return text;
	return useEllipsis
		? text.substring(0, length) + "..."
		: text.substring(0, length);
};

const showFallback = ref<boolean>(false);
</script>

<template>
	<div
		class="h-full flex flex-col bg-white dark:bg-gray-800 transition-colors duration-200"
	>
		<div class="p-4 border-r border-gray-200 dark:border-gray-700">
			<div
				class="mt-2 flex items-center space-x-2 bg-gray-100 dark:bg-gray-700 rounded-lg p-2 border-r border-gray-200 dark:border-gray-700"
			>
				<Search class="text-gray-500 dark:text-gray-400 w-5 h-5" />
				<input
					v-model="searchQuery"
					type="text"
					:placeholder='
						activeTab === "friends"
						? "搜索好友..."
						: "搜索群聊..."
					'
					@input="() => debounce(() => {}, 300)"
					class="flex-1 bg-transparent focus:outline-none text-sm dark:text-gray-200"
				/>
				<button
					class="cursor-pointer w-auth flex items-center justify-center space-x-2 p-2 bg-blue-500 hover:bg-blue-600 text-white rounded-lg transition-colors"
				>
					<UserPlus class="w-5 h-5" v-once />
				</button>
			</div>
		</div>

		<div
			class="bg-white w-auto dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700"
		>
			<div class="flex border-b border-gray-200 dark:border-gray-700">
				<button
					@click='activePage = "friendReq"'
					class="cursor-pointer text-blue-500 flex p-4 text-sm font-medium transition-colors justify-between w-full"
					v-once
				>
					<span class="text-black dark:text-white">请求信息</span>
					<ChevronRight />
				</button>
			</div>
		</div>

		<div
			class="h-full flex flex-col bg-white dark:bg-gray-800 border-r border-gray-200 dark:border-gray-700"
		>
			<div class="flex border-b border-gray-200 dark:border-gray-700">
				<button
					v-for="tab in tabs"
					:key="tab.id"
					@click="activeTab = tab.id"
					class="flex-1 p-4 text-sm font-medium transition-colors"
					:class='
						{
							"text-blue-500 border-b-2 border-blue-500":
								activeTab === tab.id,
							"text-gray-500 hover:bg-gray-50 dark:hover:bg-gray-700":
								activeTab !== tab.id,
						}
					'
				>
					{{ tab.label }}
				</button>
			</div>

			<div class="flex-1 overflow-y-auto p-4">
				<template v-if='activeTab === "friends"'>
					<div
						v-for="group in filteredFriendGroups"
						:key="group.id"
						class="mb-6"
					>
						<div class="flex items-center justify-between mb-3">
							<h3
								class="font-medium text-gray-700 dark:text-gray-300"
							>
								{{
									group
									.name
								}} ({{
									group
									.items
									.length
								}})
							</h3>
							<button
								@click="
									toggleGroup(
										group.id,
									)
								"
								class="p-1 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md"
							>
								<ChevronDown
									class="w-5 h-5 transform transition-transform dark:text-gray-500"
									:class='
										{
											"rotate-180":
												group
													.collapsed,
										}
									'
								/>
							</button>
						</div>

						<div v-show="!group.collapsed" class="space-y-2">
							<div
								v-for="friend in group.items"
								:key="friend.id"
								class="flex items-center p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg transition-colors"
							>
								<div class="relative flex-shrink-0">
									<div
										class="w-10 h-10 bg-blue-100 dark:bg-gray-700 rounded-full flex items-center justify-center"
									>
										<img
											alt="好友头像"
											v-show="!showFallback"
											:src="`/api/avatar/get/blk`"
											@error="
												showFallback =
												true
											"
											class="w-full h-full rounded-full object-cover bg-transparent"
										/>
										<span
											v-show="showFallback"
											class="text-gray-600 dark:text-gray-300 font-medium transition-opacity duration-200"
										>
											{{
												truncateText(
													friend
														.name,
													2,
													false,
												)
											}}
										</span>
									</div>
									<div
										v-if="
											friend
											.online
										"
										class="absolute bottom-0 right-0 w-3 h-3 bg-green-500 rounded-full border-2 border-white dark:border-gray-800"
									/>
								</div>
								<div class="ml-3 flex-1 min-w-0">
									<div
										class="font-medium text-gray-700 dark:text-gray-300 truncate"
									>
										{{
											truncateText(
												friend
													.name,
												6,
											)
										}}
									</div>
									<div
										class="text-sm text-gray-500 dark:text-gray-400"
									>
										{{
											friend
											.status
										}}
									</div>
								</div>
							</div>
						</div>
					</div>
				</template>

				<template v-if='activeTab === "groups"'>
					<div
						v-for="category in filteredGroupCategories"
						:key="category.id"
						class="mb-6"
					>
						<div class="flex items-center justify-between mb-3">
							<h3
								class="font-medium text-gray-700 dark:text-gray-300"
							>
								{{
									category
									.name
								}} ({{
									category
									.items
									.length
								}})
							</h3>
							<button
								@click="
									toggleCategory(
										category.id,
									)
								"
								class="p-1 hover:bg-gray-100 dark:hover:bg-gray-700 rounded-md"
							>
								<ChevronDown
									class="w-5 h-5 transform transition-transform dark:text-gray-500"
									:class='
										{
											"rotate-180":
												category
													.collapsed,
										}
									'
								/>
							</button>
						</div>

						<div v-show="!category.collapsed" class="space-y-2">
							<div
								v-for="group in category.items"
								:key="group.id"
								class="flex items-center p-3 hover:bg-gray-50 dark:hover:bg-gray-700 rounded-lg transition-colors"
							>
								<div class="relative flex-shrink-0">
									<div
										class="w-10 h-10 bg-blue-100 dark:bg-gray-700 rounded-full flex items-center justify-center"
									>
										<span
											class="text-gray-600 dark:text-gray-300 font-medium"
										>
											{{
												truncateText(
													group
														.avatar,
													2,
													false,
												)
											}}
										</span>
									</div>
								</div>

								<div class="ml-3 flex-1 min-w-0">
									<div
										class="font-medium text-gray-700 dark:text-gray-300 truncate mb-1"
									>
										{{
											truncateText(
												group
													.name,
												6,
											)
										}}
									</div>

									<div
										class="flex items-center justify-between"
									>
										<div
											class="text-sm text-gray-500 dark:text-gray-400 truncate"
										>
											{{
												truncateText(
													group
														.lastMessage ||
														"暂无消息",
													6,
												)
											}}
										</div>
										<div
											class="text-sm text-gray-400 dark:text-gray-500 ml-2 whitespace-nowrap"
										>
											{{
												group
												.members
											}}人
										</div>
									</div>
								</div>
							</div>
						</div>
					</div>
				</template>
			</div>

			<div class="p-4 border-t border-gray-200 dark:border-gray-700">
				<button
					class="w-full flex items-center justify-center space-x-2 p-2 bg-blue-500 hover:bg-blue-600 text-white rounded-lg"
				>
					<Plus class="w-5 h-5" />
					<span>{{
						activeTab === "friends"
						? "新建分组"
						: "新建分类"
					}}</span>
				</button>
			</div>
		</div>
	</div>
	<div class="w-full bg-white dark:bg-gray-800">
		<template v-if='activePage == "friendReq"'>
			<div class="h-full flex flex-col bg-white dark:bg-gray-800">
				<div class="p-4 border-b border-gray-200 dark:border-gray-700">
					<h2
						class="text-lg font-semibold text-gray-800 dark:text-gray-200"
					>
						好友申请
					</h2>
				</div>

				<div class="flex-1 overflow-y-auto">
					<div
						v-if="friendRequests.length === 0"
						class="flex flex-col items-center justify-center h-full text-gray-500 dark:text-gray-400"
					>
						<UserPlus class="w-12 h-12 mb-4 opacity-50" />
						<p>暂无好友申请</p>
					</div>

					<div
						v-else
						class="divide-y divide-gray-200 dark:divide-gray-700"
					>
						<div
							v-for="request in friendRequests"
							:key="request.id"
							class="p-4 hover:bg-gray-50 dark:hover:bg-gray-700 transition-colors"
						>
							<div class="flex items-start">
								<div class="relative mr-3 flex-shrink-0">
									<div
										class="w-12 h-12 bg-blue-100 dark:bg-gray-600 rounded-full flex items-center justify-center"
									>
										<span
											class="text-lg font-medium text-gray-600 dark:text-gray-300"
										>
											{{
												request
												.from
												.avatar
											}}
										</span>
									</div>
									<div
										v-if="
											request
											.from
											.online
										"
										class="absolute bottom-0 right-0 w-3 h-3 bg-green-500 rounded-full border-2 border-white dark:border-gray-800"
									/>
								</div>

								<div class="flex-1 min-w-0 flex flex-col">
									<div
										class="flex items-center justify-between w-full"
									>
										<h3
											class="font-medium text-gray-800 dark:text-gray-200 truncate"
										>
											{{
												request
												.from
												.name
											}}
										</h3>
										<span
											class="text-xs text-gray-500 dark:text-gray-400 whitespace-nowrap ml-2"
										>
											{{
												formatTime(
													request
														.time,
												)
											}}
										</span>
									</div>

									<p
										class="text-sm text-gray-600 dark:text-gray-300 mt-1 break-words"
									>
										{{
											request
											.message ||
											"请求添加你为好友"
										}}
									</p>

									<div
										class="flex justify-end space-x-2 mt-3"
									>
										<button
											@click='
												handleRequest(
													request
														.id,
													"accept",
												)
											'
											class="cursor-pointer px-4 py-2 bg-green-500 hover:bg-green-600 text-white rounded-md text-sm transition-colors w-20"
											v-if='
												request
												.status ==
												"pending"
											'
										>
											接受
										</button>
										<button
											@click='
												handleRequest(
													request
														.id,
													"reject",
												)
											'
											class="cursor-pointer px-4 py-2 bg-gray-200 hover:bg-gray-300 dark:bg-gray-600 dark:hover:bg-gray-500 text-gray-800 dark:text-gray-200 rounded-md text-sm transition-colors w-20"
											v-if='
												request
												.status ==
												"pending"
											'
										>
											拒绝
										</button>
										<span
											v-else
											class="py-2 px-4 rounded-lg text-sm font-medium"
											:class='
												{
													"text-green-500 bg-green-500/10":
														request
															.status ===
															"accept",
													"text-red-500 bg-red-500/10":
														request
															.status ===
															"reject",
												}
											'
										>
											{{
												request
													.status ===
													"accept"
												? "已同意"
												: "已拒绝"
											}}
										</span>
									</div>
								</div>
							</div>
						</div>
					</div>
				</div>

				<div
					class="p-3 border-t border-gray-200 dark:border-gray-700 text-sm text-gray-500 dark:text-gray-400"
				>
					共 {{ friendRequests.length }} 条申请
				</div>
			</div>
		</template>
	</div>
</template>
