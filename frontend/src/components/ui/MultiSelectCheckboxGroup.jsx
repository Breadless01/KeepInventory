export default function MultiSelectCheckboxGroup({ label, options, values, onChange }) {
    function toggle(id) {
        if (values.includes(id)) {
            onChange(values.filter((v) => v !== id));
        } else {
            onChange([...values, id]);
        }
    }

    return (
        <div className="ki-facet">
            <div className="ki-facet-label">{label}</div>
            <div className="ki-facet-options">
                {options.map((opt) => (
                    <label key={opt.ID ?? opt.id} className="ki-facet-option">
                        <input
                            type="checkbox"
                            checked={values.includes(opt.ID ?? opt.id)}
                            onChange={() => toggle(opt.ID ?? opt.id)}
                        />
                        <span className="ki-facet-option-name">{opt.Name ?? opt.name}</span>
                        <span className="ki-facet-option-count">{opt.Count ?? opt.count}</span>
                    </label>
                ))}
            </div>
        </div>
    );
}