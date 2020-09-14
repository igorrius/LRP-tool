# Parser project configuration
This folder contains project configurations in CSV format. These configurations can be overridden using the Docker volume approach.

## Configuration file format
This project uses CSV files to manage the project configuration. These files should be in the next format
(with the first line as column declaration) and using `comma` separation symbol:

scope,profession,origin_site,courses,subscription,link




Scope - file name
Specialization - professions (table)
destination - steps (table)
