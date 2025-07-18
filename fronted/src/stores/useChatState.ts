import { defineStore } from "pinia";
import { computed, ref } from "vue";

interface BaseChat {
  id: string;
  type: "private" | "group" | "temp";
  unread: number;
  lastMessage?: string;
  lastTime?: Date;
}

interface PrivateChat extends BaseChat {
  type: "private";
  friendId: string;
  friendName: string;
  friendNickName: string;
  friendAvatar: string;
}

interface GroupChat extends BaseChat {
  type: "group";
  groupId: string;
  groupName: string;
  groupAvatar: string;
}

interface TempChat extends BaseChat {
  type: "temp";
  groupId: string;
  tempSessionId: string;
  friendId: string;
  friendName: string;
  friendAvatar: string;
}

type ChatSession = PrivateChat | GroupChat | TempChat;

const useChatState = defineStore("useChatState", () => {
  const isChatting = ref(false);
  const chatSessions = ref<ChatSession[]>([]);
  const activeSessionId = ref<string | null>(null);
  const isPrivateChat = (session: ChatSession): session is PrivateChat =>
    session.type === "private";
  const isGroupChat = (session: ChatSession): session is GroupChat =>
    session.type === "group";
  const isTempChat = (session: ChatSession): session is TempChat =>
    session.type === "temp";

  const activeSession = computed(() => {
    return (
      chatSessions.value.find(
        (session) => session.id === activeSessionId.value,
      ) || null
    );
  });
  const getSessionById = (sessionId: string) => {
    return chatSessions.value.find((session) => session.id === sessionId);
  };
  const currentSessionDetail = computed(() => {
    if (!activeSession.value) return null;
    switch (activeSession.value.type) {
      case "private":
        return {
          type: "private" as const,
          friendId: activeSession.value.friendId,
          name: activeSession.value.friendNickName ||
            activeSession.value.friendName,
          avatar: activeSession.value.friendAvatar,
        };
      case "group":
        return {
          type: "group" as const,
          groupId: activeSession.value.groupId,
          name: activeSession.value.groupName,
          avatar: activeSession.value.groupAvatar,
        };
      case "temp":
        return {
          type: "temp" as const,
          groupId: activeSession.value.groupId,
          friendId: activeSession.value.friendId,
          name: activeSession.value.friendName,
          avatar: activeSession.value.friendAvatar,
        };
      default:
        return null;
    }
  });
  const privateSessions = computed(() =>
    chatSessions.value.filter(isPrivateChat)
  );
  const groupSessions = computed(() => chatSessions.value.filter(isGroupChat));
  const tempSessions = computed(() => chatSessions.value.filter(isTempChat));

  function openSession(session: ChatSession) {
    const exists = chatSessions.value.some((s) => s.id === session.id);
    if (!exists) {
      chatSessions.value.push(session);
    }
    setActiveSession(session.id);
    isChatting.value = true;
  }
  function setActiveSession(sessionId: string) {
    if (chatSessions.value.some((s) => s.id === sessionId)) {
      activeSessionId.value = sessionId;
    }
  }
  function closeSession(sessionId: string) {
    const index = chatSessions.value.findIndex((s) => s.id === sessionId);
    if (index !== -1) {
      chatSessions.value.splice(index, 1);

      if (activeSessionId.value === sessionId) {
        activeSessionId.value = chatSessions.value.length > 0
          ? chatSessions.value[0].id
          : null;
      }

      if (chatSessions.value.length === 0) {
        isChatting.value = false;
      }
    }
  }
  function updateLastMessage(sessionId: string, message: string) {
    const session = chatSessions.value.find((s) => s.id === sessionId);
    if (session) {
      session.lastMessage = message;
      session.lastTime = new Date();
      if (session.id !== activeSessionId.value) {
        session.unread += 1;
      }
    }
  }

  return {
    isChatting,
    chatSessions,
    activeSessionId,
    activeSession,
    getSessionById,
    currentSessionDetail,
    privateSessions,
    groupSessions,
    tempSessions,
    isPrivateChat,
    isGroupChat,
    isTempChat,
    openSession,
    setActiveSession,
    closeSession,
    updateLastMessage,
  };
});

export { useChatState };
export type { ChatSession, GroupChat, PrivateChat, TempChat };
