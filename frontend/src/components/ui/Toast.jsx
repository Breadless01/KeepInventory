import React, { useEffect } from "react";
import "./toast.css";

export default function Toast({
  id,
  type = "info",
  title,
  message,
  mode = "dismissible",
  onClose,
}) {
  useEffect(() => {
    if (mode === "dismissible") {
      const timer = setTimeout(() => {
        onClose(id);
      }, 5000);
      return () => clearTimeout(timer);
    }
  }, [id, mode, onClose]);

  return (
    <div className={`toast toast-${type}`}>
      <div className="toast-content">
        <strong className="toast-title">{title}</strong>
        {message && <span className="toast-message">{message}</span>}
      </div>

      {mode === "static" && (
        <button className="toast-close" onClick={() => onClose(id)}>
          ✕
        </button>
      )}
    </div>
  );
}
