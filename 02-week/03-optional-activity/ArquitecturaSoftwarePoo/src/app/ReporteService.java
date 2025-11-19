package app;

public class ReporteService {
    private EmpleadoRepository repo;

    public ReporteService(EmpleadoRepository repo) {
        this.repo = repo;
    }

    public String generarReporte(Reporte reporte) {
        return reporte.generar(repo.findAll());
    }
}