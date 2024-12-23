CREATE TABLE IF NOT EXISTS glossary
(
    id
    SERIAL
    PRIMARY
    KEY,
    term
    VARCHAR
(
    255
) NOT NULL UNIQUE,
    description TEXT NOT NULL
    );

INSERT INTO glossary (term, description)
VALUES ('WebAssembly', 'Универсальный исполняемый формат для создания более производительных веб-приложений.'),
       ('Wasm Module', 'Компилируемая единица кода в WebAssembly, содержащая функции и данные.'),
       ('Compilation', 'Процесс преобразования исходного кода в исполняемый формат.'),
       ('JIT Compilation', 'Компиляция "на лету", оптимизирующая производительность во время выполнения приложения.'),
       ('Optimization', 'Процесс улучшения производительности кода через различные методы и алгоритмы.'),
       ('Garbage Collection', 'Процесс автоматической очистки памяти, используя сборщики мусора.'),
       ('Fallback', 'Обработка ситуации, когда WebAssembly не поддерживается, через альтернативные методы.'),
       ('Threading',
        'Использование потоков для параллельного выполнения инструкций, особенно важно для многопоточных сред.'),
       ('Memory Management', 'Управление выделением и очисткой памяти для оптимизации производительности.');

INSERT INTO glossary (term, description)
VALUES ('AOT Compilation', 'Выполнение предварительной компиляции для повышения производительности.'),
       ('Polyfill', 'Код, используемый для обеспечения поддержки новых функций в старых браузерах.'),
       ('Web Workers', 'API для выполнения фоновых задач без блокировки основного потока.'),
       ('Caching', 'Хранение часто используемых данных в памяти для быстрого доступа.'),
       ('Latency', 'Задержка в обработке данных или выполнение задач в приложении.'),
       ('Bandwidth Optimization', 'Процесс сокращения потребляемого трафика для увеличения скорости работы.'),
       ('Load Balancing', 'Распределение трафика между несколькими серверами для оптимизации ресурсов.'),
       ('Profiling', 'Процесс анализа производительности для идентификации узких мест.'),
       ('Benchmarking', 'Измерение производительности системы или кода.'),
       ('Hot Module Replacement', 'Механизм замены модулей приложения без перезагрузки страницы.'),
       ('API Rate Limiting', 'Ограничение числа запросов к API за определенный период времени.'),
       ('Lazy Loading', 'Загрузка данных или модулей только по мере необходимости.'),
       ('Tree Shaking', 'Удаление неиспользуемого кода из пакета приложений.'),
       ('Service Workers',
        'Скрипты, выполняемые браузером в фоновом режиме, для улучшений работы оффлайн и кэширования.'),
       ('TypeScript', 'Язык программирования, расширяющий возможности JavaScript за счет строгой типизации.'),
       ('Babel', 'Транслятор кода, который позволяет использовать более новые возможности JavaScript.'),
       ('ESM', 'Стандартный формат модулей ECMA Script для обеспечения модульности и управления зависимостями.'),
       ('Minification', 'Процесс уменьшения размера файлов кода путем удаления пробелов и комментариев.'),
       ('Code Splitting', 'Разбиение кода на части для загрузки только необходимых ресурсов.'),
       ('Performance Budget',
        'Ограничения на ресурсы, такие как время загрузки или размер файла, для улучшения производительности.');
