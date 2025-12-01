import { useEffect, useState } from "react";
import { Modal } from "../ui/Modal.jsx";
import "./newModal.css"

import {
  CreateProjekt
} from "../../../wailsjs/go/backend/App";


export function NewProjektModal({ open, onClose, onCreated }) {
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState("");

  // Formularwerte
  const [name, setName] = useState("");
  const [kunde, setKunde] = useState("");

  useEffect(() => {
    if (!open) return;
    setError("");
  }, [open])

  async function handleSubmit(e) {
    e?.preventDefault();
    setError("");

    if (!name.trim()) {
      setError("Bitte einen Namen eingeben.");
      return;
    }

    if (!kunde.trim()) {
      setError("Bitte einen Kunden eingeben.");
      return;
    }

    try {
      setSaving(true);

      const req = {
        Name: name,
        Kunde: kunde
      };

      const created = await CreateProjekt(req);

      if (onCreated) {
        onCreated(created);
      }

      // Formular zurücksetzen
      setName("");
      setKunde("");

      onClose();
    } catch (e) {
      console.error(e);
      setError(String(e));
    } finally {
      setSaving(false);
    }
  }

  if (!open) return null;

  return (
    <Modal title="Neuer Kunde" onClose={onClose}>
      <form className="ki-form" onSubmit={handleSubmit}>
        {error && <div className="ki-error">{error}</div>}
        <div className="ki-form-group">
          <label>Name*</label>
          <input
            className="ki-input"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />
        </div>
        <div className="ki-form-group">
          <label>Kunde*</label>
          <input
            className="ki-input"
            value={kunde}
            onChange={(e) => setKunde(e.target.value)}
            required
          />
        </div>
        <div className="ki-form-actions">
          <button
            type="button"
            className="ki-btn-secondary"
            onClick={onClose}
            disabled={saving}
          >
            Abbrechen
          </button>
          <button
            type="submit"
            className="ki-btn-primary"
            disabled={saving}
          >
            {saving ? "Speichern…" : "Bauteil anlegen"}
          </button>
        </div>
      </form>
    </Modal>
  );
}