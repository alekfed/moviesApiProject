# moviesApiProject

| stack            | image size | download time | build time |
|------------------|------------|---------------|------------|
| net/http         | 6.82MB     | 0.2s          | 1.8s       |
| net/http + pq    | 7.47MB     | 0.3s          | 1.6s       |
| net/http + gpx   | 8.46MB     | 9.5s          | 2.3s       |
| gin              | 11MB       | 13.1s         | 3.2s       |
| gin + pq         | 11.3MB     | 12.5s         | 3.3s       |
| gin + gorm (pgx) | 14.6MB     | 16.7s         | 4.6s       |

(built on Apple M1)