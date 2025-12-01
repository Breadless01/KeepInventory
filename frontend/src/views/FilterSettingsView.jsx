import { useEffect, useState } from "react";
import { GetFilterConfig, SaveFilterConfig } from "../../wailsjs/go/backend/App";
import { useToasts } from "../components/ui/ToastContainer.jsx";

export default function FilterSettingsView() {
    const [config, setConfig] = useState(null);
    const [saving, setSaving] = useState(false);

    const { addToast } = useToasts();

    useEffect(() => {
        GetFilterConfig().then(setConfig);
    }, []);

    if (!config) {
        return <div>Lade Filter-Konfiguration…</div>;
    }

    function toggleField(resourceIndex, fieldIndex) {
        const next = structuredClone(config);
        const field = next.resources[resourceIndex].fields[fieldIndex];
        field.enabled = !field.enabled;
        setConfig(next);
    }

    async function handleSave() {
        setSaving(true);
        try {
            await SaveFilterConfig(config);
            addToast({
                title: "Gespeichert",
                message: "Einstellung gespeichert"
            });
        } finally {
            setSaving(false);
        }
    }

    return (
        <div className="ki-content">
            <h2 className="ki-h2">Filter-Einstellungen</h2>

            {config.resources.map((res, rIndex) => (
                <div key={res.resource} className="ki-card" style={{ marginBottom: "1rem" }}>
                    <h3 className="ki-h3">
                        <span className="ki-text-muted">{res.table}</span>
                    </h3>

                    <div className="ki-filter-fields-grid">
                        {res.fields.map((f, fIndex) => (
                            <label key={f.field} className="ki-filter-settings-row">
                                <span className="ki-filter-settings-label">
                                    {f.label}
                                </span>
                                <input
                                    type="checkbox"
                                    checked={f.enabled}
                                    onChange={() => toggleField(rIndex, fIndex)}
                                />
                            </label>
                        ))}
                    </div>
                </div>
            ))}

            <button
                className="ki-btn-primary"
                onClick={handleSave}
                disabled={saving}
            >
                {saving ? "Speichere…" : "Speichern"}
            </button>
        </div>
    );
}
