import app.*;

public class Main {
    public static void main(String[] args) {
        EmpleadoRepository repo = new EmpleadoRepository();

        repo.add(new Administrativo("A1", "Juan Pérez", 1200));
        repo.add(new Gerente("G1", "Ana Gómez", 2000, 500));

        ReporteService service = new ReporteService(repo);

        System.out.println("=== Reporte en texto ===");
        System.out.println(service.generarReporte(new ReporteTexto()));

        System.out.println("=== Reporte en CSV ===");
        System.out.println(service.generarReporte(new ReporteCSV()));
    }
}
