importScripts("https://www.gstatic.com/firebasejs/10.12.2/firebase-app-compat.js");
importScripts("https://www.gstatic.com/firebasejs/10.12.2/firebase-messaging-compat.js");

firebase.initializeApp({
  apiKey: "AIzaSyBwIdUiaxKmrN6dwoMxvLzkuOX9WJIcIf4",
  authDomain: "skproject-9403d.firebaseapp.com",
  projectId: "skproject-9403d",
  messagingSenderId: "656977741555",
  appId: "1:656977741555:web:6f83f02cf8de0a3538d986"
});

const messaging = firebase.messaging();

// ✅ Background message handler
messaging.onBackgroundMessage((payload) => {
  console.log("Received background message ", payload);
  const notificationTitle = payload.notification.title;
  const notificationOptions = {
    body: payload.notification.body,
    icon: "/icon.png"
  };
  self.registration.showNotification(notificationTitle, notificationOptions);
});
