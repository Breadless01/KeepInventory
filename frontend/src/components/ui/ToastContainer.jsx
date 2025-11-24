import React, { createContext, useContext, useState, useCallback } from "react";
import Toast from "./Toast";
import "./toast.css";

const ToastContext = createContext(null);

export function useToasts() {
  return useContext(ToastContext);
}

export function ToastProvider({ children }) {
  const [toasts, setToasts] = useState([]);

  const addToast = useCallback(({ title, message, type = "info", mode = "dismissible" }) => {
    const id = crypto.randomUUID();
    setToasts((prev) => [...prev, { id, title, message, type, mode }]);
  }, []);

  const removeToast = useCallback((id) => {
    setToasts((prev) => prev.filter((t) => t.id !== id));
  }, []);

  return (
    <ToastContext.Provider value={{ addToast }}>
      {children}

      <div className="toast-container">
        {toasts.map((t) => (
          <Toast key={t.id} {...t} onClose={removeToast} />
        ))}
      </div>
    </ToastContext.Provider>
  );
}
