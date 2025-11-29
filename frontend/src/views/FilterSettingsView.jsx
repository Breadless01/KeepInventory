import { useEffect, useState } from "react";
import { GetFilterConfig, SaveFilterConfig } from "../../wailsjs/go/backend/App";

export default function FilterSettingsView() {
    const [config, setConfig] = useState(null);
    const [saving, setSaving] = useState(false);

    useEffect(() => {
        GetFilterConfig().then(setConfig);
    }, []);

    if (!config) {
        return <div>Lade Filter-Konfiguration…</div>;
    }

    function toggleField(resourceIndex, fieldIndex) {
        const next = structuredClone(config);
        const field = next.Resources[resourceIndex].Fields[fieldIndex];
        field.Enabled = !field.Enabled;
        setConfig(next);
    }

    async function handleSave() {
        setSaving(true);
        try {
            await SaveFilterConfig(config);
            // optional Toast
        } finally {
            setSaving(false);
        }
    }

    return (
        <div className="ki-content">
            <h2 className="ki-h2">Filter-Einstellungen</h2>

            {config.Resources.map((res, rIndex) => (
                <div key={res.Resource} className="ki-card" style={{ marginBottom: "1rem" }}>
                    <h3 className="ki-h3">
                        Ressource: {res.Resource} <span className="ki-text-muted">({res.Table})</span>
                    </h3>

                    <div className="ki-filter-fields-grid">
                        {res.Fields.map((f, fIndex) => (
                            <label key={f.Field} className="ki-filter-settings-row">
                                <input
                                    type="checkbox"
                                    checked={f.Enabled}
                                    onChange={() => toggleField(rIndex, fIndex)}
                                />
                                <span className="ki-filter-settings-label">
                  {f.Label} <span className="ki-text-soft">({f.Field})</span>
                </span>
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
