package app;
import java.util.ArrayList;
import java.util.List;

public class EmpleadoRepository {
    private List<Empleado> storage = new ArrayList<>();

    public void add(Empleado e) { storage.add(e); }
    public List<Empleado> findAll() { return new ArrayList<>(storage); }
}