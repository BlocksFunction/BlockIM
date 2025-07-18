import ReactiveDB from "@/lib/tool/ReactiveDB.ts";
import { defineStore } from "pinia";
import { ref } from "vue";

enum ReqStatus {
  pending,
  success,
  rejected,
  failed,
}

interface ReqFrom {
  id: number;
  name: string;
  avatar: string;
  online: boolean;
}

interface FriendReq {
  id: number;
  from: ReqFrom;
  message: string;
  time: Date;
  status: ReqStatus;
}

const useFriendReq = defineStore("friend/useFriendReq", () => {
  const db = new ReactiveDB<FriendReq>({
    dbName: "friend",
    storeName: "friendReq",
    useShallowRef: true,
    autoSaveDelay: 500,
  });
  const dbInitialized = ref(false);

  async function init() {
    try {
      await db.initialize();
      dbInitialized.value = true;
    } catch (error) {
      throw error;
    }
  }
});
